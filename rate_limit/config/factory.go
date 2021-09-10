package config

import "fmt"

type Format string

func NewConfigParser(fromat string) (Interface, error) {
	if 0 == len(fromat) {
		return nil, fmt.Errorf("NewConfigParser: format must be one of [yaml,json]")
	}

	switch fromat {
	case "yaml":
		return NewParseYaml(), nil
	case "json":
		return NewParseJson(), nil
	default:
		return nil, fmt.Errorf("NewConfigParser: format must be one of [yaml,json]")
	}
}
