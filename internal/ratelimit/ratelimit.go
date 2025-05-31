package ratelimit

import (
	"context"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type RateLimit struct {
	refill       float64 // per second
	bucketsize   int
	cleanupAfter time.Duration

	mu       sync.RWMutex
	visitors map[string]*visitor
}

type visitor struct {
	limit *rate.Limiter

	mu       sync.RWMutex
	lastSeen time.Time
}

func (v *visitor) Allow() bool {
	v.mu.Lock()
	v.lastSeen = time.Now()
	v.mu.Unlock()
	return v.limit.Allow()
}

func (v *visitor) getLastSeen() time.Time {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.lastSeen
}

func NewRateLimitContext(ctx context.Context, refill float64, bucketsize int) *RateLimit {
	cleanupAfter := max(time.Duration((float64(bucketsize)/refill)*3)*time.Second, time.Minute)
	rl := &RateLimit{
		refill:       refill,
		bucketsize:   bucketsize,
		cleanupAfter: cleanupAfter,
		mu:           sync.RWMutex{},
		visitors:     make(map[string]*visitor),
	}
	go rl.cleanupVisitors(ctx)
	return rl
}

func (rl *RateLimit) GetVisitor(ip string) *visitor {
	rl.mu.RLock()
	v, found := rl.visitors[ip]
	rl.mu.RUnlock()
	if !found {
		v = &visitor{
			limit:    rate.NewLimiter(rate.Limit(rl.refill), rl.bucketsize),
			lastSeen: time.Now(),
		}
		rl.mu.Lock()
		rl.visitors[ip] = v
		rl.mu.Unlock()
	}
	return v
}

func (rl *RateLimit) cleanupVisitors(ctx context.Context) {
	tick := time.NewTicker(time.Minute)
	defer tick.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			rl.mu.Lock()
			for ip, v := range rl.visitors {
				if time.Since(v.getLastSeen()) > rl.cleanupAfter {
					delete(rl.visitors, ip)
				}
			}
			rl.mu.Unlock()
		}
	}
}
