package main

import (
	"os"

	. "github.com/dave/jennifer/jen"
	"github.com/gertd/go-pluralize"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	titleCaser = cases.Title(language.English)
	lower      = cases.Lower(language.English)
	pul        = pluralize.NewClient()
)

func main() {
	f := NewFilePath("main")
	f.HeaderComment("Code generated by genum DO NOT EDIT")

	t := EnumType{
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
	}

	f.Add(t.toCode()...)

	create, _ := os.Create("countries.go")
	create.WriteString(f.GoString())
}
