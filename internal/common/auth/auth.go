package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/ivanwang123/roadmap/internal/common/cookie"
)

var userIDCtxKey = "user_id"

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// valid user cookie = add user id to ctx
			// expired user cookie = get refresh cookie & create new user & refresh cookie & add user id to ctx
			// expired user & refresh cookie = noauth continue
			// no user cookie = get refresh cookie & create new user & refresh cookie & add user id to ctx
			// no user or refresh cookie = noauth continue
			var userId int
			userCookie, err := r.Cookie("user")

			if err != nil {
				userId, err = RefreshToken(r)
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}
			} else {

				userId, err = ParseToken(userCookie.Value)
				if err != nil {
					userId, err = RefreshToken(r)
					userCookie.MaxAge = -1
					if err != nil {
						next.ServeHTTP(w, r)
						return
					}
				}
			}

			if userId > 0 {
				ctx := context.WithValue(r.Context(), userIDCtxKey, userId)
				r = r.WithContext(ctx)
			}
			next.ServeHTTP(w, r)
		})
	}
}

func RefreshToken(r *http.Request) (int, error) {
	refreshCookie, err := r.Cookie("refresh")
	if err != nil {
		return -1, err
	}

	userId, err := ParseToken(refreshCookie.Value)
	if err != nil {
		refreshCookie.MaxAge = -1
		return -1, err
	}

	refreshCookie.MaxAge = -1
	userToken, err := GenerateToken(userId, time.Hour*24)
	if err != nil {
		return -1, err
	}
	refreshToken, err := GenerateToken(userId, time.Hour*24*7)
	if err != nil {
		return -1, err
	}

	cookie.ForContext(r.Context()).SetCookie("user", userToken, time.Hour*24)
	cookie.ForContext(r.Context()).SetCookie("refresh", refreshToken, time.Hour*24*7)

	return userId, nil
}

func GetCurrentUser(ctx context.Context) (int, error) {
	raw := ctx.Value(userIDCtxKey)
	if raw != nil {
		return raw.(int), nil
	} else {
		return -1, errors.New("User is unauthenticated")
	}
}

// TODO: Hide key
var secretKey = []byte("P98Cx7eHN57Wx82p")

func GenerateToken(userId int, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ParseToken(tokenStr string) (int, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) <= time.Now().Unix() {
				return -1, errors.New("Token expired")
			}
		} else {
			return -1, errors.New("Expiration does not exist")
		}

		if userId, ok := claims["userId"].(float64); ok {
			return int(userId), nil
		} else {
			return -1, errors.New("User id does not exist")
		}
	} else {
		return -1, err
	}
}
