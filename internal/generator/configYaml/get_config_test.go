package configYaml

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetConfig(t *testing.T) {
	t.Parallel()

	t.Run("read_file_error", func(t *testing.T) {
		config, err := GetConfig(gofakeit.Name())
		require.Nil(t, config)
		require.NotNil(t, err)
	})

	t.Run("unmarshal_config_file_error", func(t *testing.T) {
		config, err := GetConfig("./testCases/config_cases/no_version.yaml")
		require.Nil(t, config)
		require.NotNil(t, err)
	})

	t.Run("correct_config_response", func(t *testing.T) {
		config, err := GetConfig("./testCases/config_cases/no_flag_descriptions.yaml")
		require.NotNil(t, config)
		require.Nil(t, err)
	})
}
