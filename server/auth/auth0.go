package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/go-chi/chi"
)

const emailCtxKey = "email"

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

func HandleLogin() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		fmt.Println("LOGIN CODE", code)

		data := url.Values{}
		data.Set("grant_type", "authorization_code")
		data.Set("client_id", "2El7qcsqxaGzhb6ys9SME8Ofxdcvst34")
		data.Set("client_secret", "bdT8rWiG4OMRilWdAH9POZBnR-Kn7URJ1HG5T1VYYxdhoWmUOurxMZFGNDtEfi9m")
		data.Set("code", code)
		data.Set("redirect_uri", "http://localhost:3000")

		fmt.Println("DATA", data)

		req, err := http.NewRequest("POST", "https://dev-jkn4emz6.us.auth0.com/oauth/token", strings.NewReader(data.Encode()))
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			w.Write([]byte("Fail to login"))
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				w.Write([]byte("Fail to read body"))
			}
			bodyString := string(bodyBytes)
			fmt.Println("BODY", bodyString)
			cookie := &http.Cookie{
				Name:     "accesstoken",
				Value:    bodyString,
				Path:     "/",
				Expires:  time.Now().Add(time.Hour * 24),
				MaxAge:   86400,
				HttpOnly: true,
				Secure:   true,
			}
			http.SetCookie(w, cookie)
			w.Write([]byte("Success"))
		} else {
			w.Write([]byte("Fail"))
		}

	})
}

func Middleware() func(http.Handler) http.Handler {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: validationKeyGetter(),
		SigningMethod:       jwt.SigningMethodRS256,
		CredentialsOptional: true,
	})

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := jwtMiddleware.CheckJWT(w, r)
			if err != nil {
				fmt.Println("JWT MIDDLEWARE ERROR", err)
				// http.Error(w, "Invalid access token", http.StatusInternalServerError)
				// next.ServeHTTP(w, r)
				return
			}

			user := r.Context().Value("user")
			fmt.Println("USER", user)

			if user != nil {
				fmt.Println("AUTHORIZED!")
				claims := user.(*jwt.Token).Claims.(jwt.MapClaims)
				email := claims["https://roadmap.com/email"].(string)

				ctx := context.WithValue(r.Context(), emailCtxKey, email)

				r = r.WithContext(ctx)
			}
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) string {
	raw, _ := ctx.Value(emailCtxKey).(string)
	return raw
}

func validationKeyGetter() jwt.Keyfunc {
	fmt.Println("VALIDATION KEY GETTER")
	return func(token *jwt.Token) (interface{}, error) {
		// Verify 'aud' claim
		aud := "https://dev-jkn4emz6.us.auth0.com/api/v2/"
		checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
		if !checkAud {
			return token, errors.New("Invalid audience")
		}
		// Verify 'iss' claim
		iss := "https://dev-jkn4emz6.us.auth0.com/"
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return token, errors.New("Invalid issuer")
		}

		cert, err := getPemCert(token)
		if err != nil {
			panic(err.Error())
		}

		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	}
}

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get("https://dev-jkn4emz6.us.auth0.com/.well-known/jwks.json")

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}
