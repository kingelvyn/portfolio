package chat

import (
	"net/http"
	"sync"
	"time"
)

// rateLimiter is a simple per-IP rate limiter.
// Each IP gets a bucket of tokens that refills over time.
type rateLimiter struct {
	mu      sync.Mutex
	buckets map[string]*bucket
	rate    time.Duration // how often one token is added
	max     int           // max tokens in bucket
}

type bucket struct {
	tokens   int
	lastSeen time.Time
}

// newRateLimiter creates a limiter allowing max requests per window.
// Example: newRateLimiter(10, time.Minute) → 10 req/min per IP.
func newRateLimiter(maxRequests int, window time.Duration) *rateLimiter {
	rl := &rateLimiter{
		buckets: make(map[string]*bucket),
		rate:    window / time.Duration(maxRequests),
		max:     maxRequests,
	}
	// Background cleanup of stale IPs
	go rl.cleanup()
	return rl
}

func (rl *rateLimiter) allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	b, ok := rl.buckets[ip]
	if !ok {
		rl.buckets[ip] = &bucket{tokens: rl.max - 1, lastSeen: now}
		return true
	}

	// Refill tokens based on elapsed time
	elapsed := now.Sub(b.lastSeen)
	refill := int(elapsed / rl.rate)
	if refill > 0 {
		b.tokens += refill
		if b.tokens > rl.max {
			b.tokens = rl.max
		}
		b.lastSeen = now
	}

	if b.tokens <= 0 {
		return false
	}
	b.tokens--
	return true
}

// cleanup removes IPs not seen in the last 10 minutes.
func (rl *rateLimiter) cleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	for range ticker.C {
		rl.mu.Lock()
		for ip, b := range rl.buckets {
			if time.Since(b.lastSeen) > 10*time.Minute {
				delete(rl.buckets, ip)
			}
		}
		rl.mu.Unlock()
	}
}

// Middleware wraps an http.Handler with rate limiting.
func (rl *rateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := realIP(r)
		if !rl.allow(ip) {
			http.Error(w, `{"error":"rate limit exceeded, please slow down"}`, http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// realIP extracts the client IP, respecting common proxy headers.
func realIP(r *http.Request) string {
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	}
	return r.RemoteAddr
}