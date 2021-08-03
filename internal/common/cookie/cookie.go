package cookie

import (
	"context"
	"net/http"
	"time"
)

const cookieAccessCtxKey = "cookie_access"

type CookieAccess struct {
	writer http.ResponseWriter
}

func (c *CookieAccess) SetCookie(name, value string, duration time.Duration) {
	http.SetCookie(c.writer, &http.Cookie{
		Name:     name,
		Value:    value,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		Expires:  time.Now().Add(duration),
		MaxAge:   int(duration.Seconds()),
	})
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookieAccess := CookieAccess{
				writer: w,
			}
			ctx := context.WithValue(r.Context(), cookieAccessCtxKey, &cookieAccess)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *CookieAccess {
	raw, _ := ctx.Value(cookieAccessCtxKey).(*CookieAccess)
	return raw
}
