package algo

import (
	"fmt"

	"github.com/practice/rate_limit/config"
)

func New(algoType string, limit *config.Limit) (Interface, error) {
	if 0 == len(algoType) {
		return nil, fmt.Errorf("NewConfigParser: format must be one of [yaml,json]")
	}

	switch algoType {
	case "RollingWindow":
		return NewRollingWindow(limit), nil
	default:
		return nil, fmt.Errorf("NewConfigParser: format must be one of [yaml,json]")
	}
}
