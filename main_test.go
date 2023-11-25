package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Generate(t *testing.T) {
	t.Run("should generate code from the enum struct", func(t *testing.T) {
		source := "testdata/countries/country.go"
		types, err := parseSource(source)
		require.NoError(t, err)

		file, err := os.ReadFile("testdata/countries/expected_country_enums.go")
		require.NoError(t, err)

		s := generate(types, "countries")
		require.Equal(t, string(file), s)
	})
}
