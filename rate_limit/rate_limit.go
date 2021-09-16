package ratelimit

import (
	"fmt"

	"github.com/practice/rate_limit/algo"
	"github.com/practice/rate_limit/config"
)

type RateLimiter struct {
	configParser config.Interface

	// key is the app id
	limiter map[string]limiterAlgo
}

// key is the path
type limiterAlgo map[string]algo.Interface

func NewRateLimiter(configType, configFilePath, algoType string) (*RateLimiter, error) {

	rl := &RateLimiter{
		limiter: make(map[string]limiterAlgo),
	}

	// set configParser
	if err := rl.setConfigParser(configType); nil != err {
		return nil, fmt.Errorf("NewRateLimiter: %s", err)
	}

	// create algorithm for every path of apps
	if err := rl.setLimter(configFilePath, algoType); nil != err {
		return nil, fmt.Errorf("NewRateLimiter: %s", err)
	}

	return rl, nil
}

func (rl *RateLimiter) setConfigParser(configType string) error {
	configParser, err := config.New(configType)
	if nil != err {
		return fmt.Errorf("NewRateLimiter: %s", err)
	}

	rl.configParser = configParser
	return nil
}

func (rl *RateLimiter) setLimter(configFilePath, algoType string) error {
	configs, err := rl.configParser.Parse(configFilePath)
	if nil != err {
		return fmt.Errorf("setLimter: %s", err)
	}

	for _, rule := range configs.Rules {
		appId := rule.AppId
		if 0 == len(appId) {
			continue
		}

		if nil == rl.limiter[appId] {
			rl.limiter[appId] = make(limiterAlgo)
		}

		for _, limit := range rule.Limits {
			algorithm, err := algo.New(algoType, limit)
			if nil != err {
				return fmt.Errorf("setLimter: %s", err)
			}
			rl.limiter[appId][limit.Path] = algorithm
		}
	}

	return nil
}

func (rl *RateLimiter) Limit(appId, path string) bool {
	limiter := rl.limiter[appId][path]
	if nil == limiter {
		return false
	}

	return limiter.Limit()
}
