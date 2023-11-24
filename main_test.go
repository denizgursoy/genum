package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Generate(t *testing.T) {
	t.Run("should generate code from the enum struct", func(t *testing.T) {
		types := []EnumType{
			{
				name: "country",
				fields: FieldTypes{
					{
						Name: "name",
						Type: "string",
					},
					{
						Name: "shorthand",
						Type: "string",
					},
					{
						Name: "continent",
						Type: "string",
					},
				},
				enumValues: []EnumValue{
					{
						Name: "TR",
						fields: FieldValues{
							{
								Name:  "name",
								Value: "Turkey",
							},
						},
					},
					{
						Name: "NL",
						fields: FieldValues{
							{
								Name:  "name",
								Value: "The Netherlands",
							},
						},
					},
					{
						Name: "JP",
						fields: FieldValues{
							{
								Name:  "name",
								Value: "Japan",
							},
						},
					},
				},
			},
		}
		file, err := os.ReadFile("testdata/countries/expected_country_enums.go")
		require.NoError(t, err)
		s := generate(types, "countries")
		require.Equal(t, string(file), s)
	})
}
