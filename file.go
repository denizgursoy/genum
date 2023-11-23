package main

import (
	"github.com/gertd/go-pluralize"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var titleCaser = cases.Title(language.English)
var lower = cases.Lower(language.English)
var pul = pluralize.NewClient()
