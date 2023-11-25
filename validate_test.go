package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_validateSource(t *testing.T) {
	t.Run("should return error if source is not valid", func(t *testing.T) {
		require.Error(t, validateSource(nil))

		source := ""
		require.Error(t, validateSource(&source))

		source = "a.java"
		require.Error(t, validateSource(&source))

		source = " testdata/countries/country.go "
		require.NoError(t, validateSource(&source))
		require.Equal(t, source, strings.TrimSpace(source))

		source = "non_present_file.go"
		require.Error(t, validateSource(&source))
	})
}
