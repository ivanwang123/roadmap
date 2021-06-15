package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/ivanwang123/roadmap/server/stores"
)

// TODO: Make secret
var secretKey = []byte("secret_auth_token")

const userCtxKey = "user"

func GenerateToken(id int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return "", nil
	}

	return tokenStr, nil
}

func ParseToken(tokenStr string) (int, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["id"].(int)
		return id, nil
	} else {
		return -1, err
	}
}

func Middleware(store *stores.Store) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr := r.Header.Get("Authorization")
			fmt.Println("TOKEN STR", tokenStr)

			if tokenStr == "" {
				next.ServeHTTP(w, r)
				return
			}

			id, err := ParseToken(tokenStr)
			fmt.Println("PARSE TOKEN", id, err)
			if err != nil {
				// "Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjM1OTczMjAsInVzZXJuYW1lIjoidXNlcjIifQ.6GU0kfCyP7KZDFwiBWfdncpawEcHa774i3Y4YC9JX7Q"
				http.Error(w, "Invalid token", http.StatusForbidden)
				// next.ServeHTTP(w, r)
				return
			}

			user, err := store.UserStore.GetById(id)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), userCtxKey, user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}
