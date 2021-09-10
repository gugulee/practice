package config

import (
	"encoding/json"
)

type Rules struct {
	Config []Config `json:"configs" yaml:"configs"`
}

type Config struct {
	// the app identify
	AppId string `json:"appId" yaml:"appId"`

	// the limit for an api interface
	Limits []Limit `json:"limits" yaml:"limits"`
}

type Limit struct {
	// the url path, e.g. /api/v1/names
	Path string `json:"path" yaml:"path"`
	// the max request limits at interval(e.g. 1s, 60s, 3600s)
	Limit int `json:"limit" yaml:"limit"`

	// the intervals
	Interval int `json:"intervals" yaml:"interval"`
}

// set default value
func (l *Limit) UnmarshalJSON(data []byte) error {
	type alias Limit
	limit := &alias{
		Interval: 1,
	}

	if err := json.Unmarshal(data, limit); nil != err {
		return err
	}

	*l = Limit(*limit)
	return nil
}

// set default value
func (l *Limit) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type alias Limit
	limit := &alias{
		Interval: 1,
	}

	if err := unmarshal(&limit); err != nil {
		return err
	}

	*l = Limit(*limit)

	return nil
}
