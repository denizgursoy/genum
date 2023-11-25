package main

import (
	"github.com/dave/jennifer/jen"
)

type (
	EnumType struct {
		name       string
		fields     FieldTypes
		enumValues EnumValues
	}

	EnumValue struct {
		Name   string
		fields FieldValues
	}

	FieldTypes  []FieldType
	FieldValues []FieldValue
	EnumValues  []EnumValue

	FieldType struct {
		Name string
		Type string
	}

	FieldValue struct {
		Name  string
		Value any
	}
)

func (r EnumType) toCode() []jen.Code {
	structName := MakeFirstLetterUpperCase(r.name)
	pluralStructName := MakePlural(structName)
	add := jen.
		Comment("type definition").Line().
		Type().Id(structName).Struct(r.fields.toCode()...).Line().
		Comment(pluralStructName).Line().
		Var().Defs(r.enumValues.toCode(structName)...).Line().
		Add(r.fields.toGetterCodes(structName)...).
		Add(getAllFunction(structName, r.enumValues)...)

	return *add
}

func getAllFunction(structName string, vals EnumValues) []jen.Code {
	all := make([]jen.Code, 0)
	count := len(vals)
	for _, val := range vals {
		a := jen.Id(val.Name)
		all = append(all, a)
	}
	a := jen.Func().Id("All" + MakePlural(structName)).
		Params().Index(jen.Lit(count)).Id(structName).
		Block(jen.Return(jen.Index(jen.Lit(count)).Id(structName).Values(all...)))

	return *a
}

func (f FieldTypes) toCode() []jen.Code {
	statements := make([]jen.Code, 0)
	for _, fieldType := range f {
		id := jen.Id(MakeFirstLetterLowerCase(fieldType.Name)).Id(fieldType.Type)
		statements = append(statements, id)
	}

	return statements
}

func (f FieldTypes) toGetterCodes(structName string) []jen.Code {
	statements := make([]jen.Code, 0)
	for _, fieldType := range f {

		id := jen.Func().Params(jen.Id("c").Id(structName)).Id(MakeFirstLetterUpperCase(fieldType.Name)).
			Params().Id(fieldType.Type).Block(
			jen.Return(jen.Id("c").Dot(fieldType.Name)),
		).Line().Line()

		statements = append(statements, id)
	}

	return statements
}

func (r EnumValues) toCode(structName string) []jen.Code {
	statements := make([]jen.Code, 0)

	for _, value := range r {
		block := jen.Line().
			Id(value.Name).Op("=").Id(structName).
			Block(value.fields.toCode()...).Line()

		statements = append(statements, block)
	}

	return statements
}

func (r FieldValues) toCode() []jen.Code {
	statements := make([]jen.Code, 0)

	for _, value := range r {
		code := jen.Id(MakeFirstLetterLowerCase(value.Name)).Id(":").Lit(value.Value).Id(",")
		statements = append(statements, code)
	}

	return statements
}
