package main

import (
	"strings"

	"github.com/gertd/go-pluralize"
)

var pul = pluralize.NewClient()

func MakePlural(name string) string {
	return pul.Plural(name)
}

func MakeFirstLetterLowerCase(name string) string {
	if len(name) < 1 {
		return name
	}

	return strings.ToLower(name[:1]) + name[1:]
}

func MakeFirstLetterUpperCase(name string) string {
	if len(name) < 1 {
		return name
	}

	return strings.ToUpper(name[:1]) + name[1:]
}

func GetFirstLetterInLowerCase(name string) string {
	if len(name) < 1 {
		return name
	}

	return strings.ToLower(name[:1])
}
