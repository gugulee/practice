package config_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/practice/rate_limit/config"
)

func Test_NewConfigParser(t *testing.T) {
	t.Run("config format is yaml", func(t *testing.T) {
		parser, err := config.NewConfigParser("yaml")
		require.NoError(t, err)
		require.IsType(t, (*config.ParseYaml)(nil), parser)
	})

	t.Run("config format is json", func(t *testing.T) {
		parser, err := config.NewConfigParser("json")
		require.NoError(t, err)
		require.IsType(t, (*config.ParseJson)(nil), parser)
	})

	t.Run("config format is empty", func(t *testing.T) {
		_, err := config.NewConfigParser("")
		require.EqualError(t, err, "NewConfigParser: format must be one of [yaml,json]")
	})

	t.Run("config format is others", func(t *testing.T) {
		_, err := config.NewConfigParser("others")
		require.EqualError(t, err, "NewConfigParser: format must be one of [yaml,json]")
	})
}
