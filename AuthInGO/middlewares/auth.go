package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	env "AuthInGo/config/env"

	"github.com/golang-jwt/jwt/v5"
)

// Define custom types for context keys
type contextKey string

const (
	userIDKey contextKey = "userId"
	emailKey  contextKey = "email"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if token == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {

			return []byte(env.GetString("JWT_SECRET", "TOKEN")), nil
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		userId, ok := claims["id"].(float64)
		email, okEmail := claims["email"].(string)

		if !ok || !okEmail {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		fmt.Println("Authenticated user ID:", int(userId), "Email:", email)

		ctx := r.Context()
		ctx = context.WithValue(ctx, userIDKey, strconv.FormatFloat(userId, 'f', -1, 64))
		ctx = context.WithValue(ctx, emailKey, email)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
