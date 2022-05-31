package argParserConfig

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestApplicationDescriptionGetters(t *testing.T) {
	t.Parallel()

	t.Run("simple", func(t *testing.T) {
		pointer := &ApplicationDescription{
			AppName:             gofakeit.Name(),
			NameHelpInfo:        gofakeit.Name(),
			DescriptionHelpInfo: gofakeit.Name(),
		}

		require.Equal(t, pointer.AppName, pointer.GetAppName())
		require.Equal(t, pointer.NameHelpInfo, pointer.GetNameHelpInfo())
		require.Equal(t, pointer.DescriptionHelpInfo, pointer.GetDescriptionHelpInfo())
	})
}
