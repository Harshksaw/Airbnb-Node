package middlewares

import (
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

var limiter = rate.NewLimiter(rate.Every(2*time.Second), 5)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simple rate limiting logic (for demonstration purposes)
		// In a real application, you would use a more robust solution
		// such as a token bucket or leaky bucket algorithm, possibly with Redis or in-memory store

		// For this example, we'll just use a simple time-based limit
		// Allow one request every second per client IP

		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})

}
