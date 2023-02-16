package middleware

import (
	"monitor/auth"
	Err "monitor/error"
	"monitor/view"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
)

func SetContentType(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		head := w.Header()
		head.Add("content-type", "application/json")
		h.ServeHTTP(w, r)
	})
}

func CheckContentType(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("content-type") != "application/json" {
			view.ErrHandle(Err.ErrBadRequestBadFields).ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func Authorization(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] == nil {
			view.ErrHandle(Err.ErrNoPermission).ServeHTTP(w, r)
			return
		}
		token, err := jwt.Parse(r.Header["Authorization"][0], func(t *jwt.Token) (interface{}, error) {
			return auth.SecretKey, nil
		})
		if err != nil || !token.Valid {
			view.ErrHandle(Err.ErrNoPermission).ServeHTTP(w, r)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			view.ErrHandle(Err.ErrNoPermission).ServeHTTP(w, r)
			return
		}
		user_id := int64(claims["id"].(float64))
		r.Header.Add("User_id", strconv.FormatInt(user_id, 10))
		h.ServeHTTP(w, r)
	})
}
