package test_arg_value

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/argtools/pkg/parsedData"
	"math/rand"
	"testing"
)

func TestArgValueToString(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName string
		argValue parsedData.ArgValue
	}{
		{
			caseName: "empty_string",
			argValue: "",
		},
		{
			caseName: "some_string",
			argValue: parsedData.ArgValue(getRandString()),
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			res, err := td.argValue.ToString()

			require.Nil(t, err)
			require.Equal(t, string(td.argValue), res)
		})
	}
}

func getRandString() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, gofakeit.Uint8())
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
