package ratelimit

import "github.com/practice/rate_limit/config"

type RateLimiter struct {
	configParser config.Interface
}

func NewRateLimiter(configParser config.Interface) *RateLimiter {
	return &RateLimiter{
		configParser: configParser,
	}
}

func (r *RateLimiter) Limit(appId, path string) bool {
	
	return true
}
