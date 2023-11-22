package main

import (
	"os"

	. "github.com/dave/jennifer/jen"
)

func main() {
	f := NewFile("countries")
	f.Type().Id("Country").Struct(
		Id("name").String())

	f.Var().Id("TR").Op("=").Id("Country").Block(
		Id("name").Id(":").Lit("a").Id(","),
	)
	f.Func().Params(Id("c").Id("Country")).Id("Name").
		Params().Id("string").Block(
		Return(Id("c").Dot("name")),
	)

	f.Func().Id("AllCountries").
		Params().Index().Id("Country").
		Block(
			Return(Index().Id("Country").Block(Id("TR").Id(","))),
		)
	create, _ := os.Create("testdata/countries.go")
	create.WriteString(f.GoString())
}
