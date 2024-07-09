package middleware

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
	"watchman/schema"

	"golang.org/x/time/rate"
)

// Defining visitor struct for use in array of unique visitors
type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// The visitors map is used to keep track of the visitors based on their IP addresses
// The mu mutex is used to protect the visitors map from concurrent reads and writes
var (
	visitors = make(map[string]*visitor)
	mu       sync.Mutex
)

// Get the visitor from the visitors map based on the IP address
func getVisitor(ip string, config schema.ConfigType) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	v, exists := visitors[ip]
	if !exists {
		limiter := rate.NewLimiter(1, config.RateLimitRequestsPerSecond)
		// Include the current time when creating a new visitor
		visitors[ip] = &visitor{limiter, time.Now()}
		return limiter
	}

	// Update the last seen time for the visitor
	v.lastSeen = time.Now()
	return v.limiter
}

// Delete the visitor if it was last seen over 3 minutes ago
func CleanupVisitors() {
	for {
		time.Sleep(time.Minute)

		mu.Lock()
		for ip, v := range visitors {
			if time.Since(v.lastSeen) > 3*time.Minute {
				delete(visitors, ip)
			}
		}
		mu.Unlock()
	}
}

func Ratelimit(config schema.ConfigType, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		limiter := getVisitor(ip, config)
		if !limiter.Allow() {
			response := schema.Response_Type{
				Status:    "ERROR",
				Message:   "You made too many requests",
				RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
			}

			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(response)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		next.ServeHTTP(w, r)
	})
}
