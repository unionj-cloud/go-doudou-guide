package httpsrv

import (
	"github.com/unionj-cloud/go-doudou/ratelimit"
	"net/http"
	"strings"
)

// RateLimit limit rate
func RateLimit(store *ratelimit.MemoryStore) func(inner http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := r.RemoteAddr[:strings.LastIndex(r.RemoteAddr, ":")]
			if !store.GetLimiter(key).Allow() {
				http.Error(w, "too many requests", http.StatusTooManyRequests)
				return
			}
			inner.ServeHTTP(w, r)
		})
	}
}
