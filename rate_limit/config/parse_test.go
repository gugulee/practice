package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var mockConfig = &Rules{
	Config: []Config{
		{
			AppId: "app-1",
			Limits: []Limit{
				{
					Path:      "/v1/user",
					Limit:     100,
					Interval: 1,
				},
				{
					Path:      "/v1/order",
					Limit:     50,
					Interval: 1,
				},
			},
		},
		{
			AppId: "app-2",
			Limits: []Limit{
				{
					Path:      "/v1/user",
					Limit:     50,
					Interval: 1,
				},
				{
					Path:      "/v1/order",
					Limit:     50,
					Interval: 1,
				},
			},
		},
	},
}

func Test_ParseJson_Parse(t *testing.T) {
	p := NewParseJson()
	config, err := p.Parse("./config.json")
	require.NoError(t, err)
	require.Equal(t, mockConfig, config)
}

func Test_ParseYaml_Parse(t *testing.T) {
	p := NewParseYaml()
	config, err := p.Parse("./config.yaml")
	require.NoError(t, err)
	require.Equal(t, mockConfig, config)
}
