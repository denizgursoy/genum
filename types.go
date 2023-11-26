package main

import (
	"fmt"

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

const (
	OrdinalField      = "_ordinal"
	NameField         = "_name"
	DefaultFieldCount = 2
)

var (
	defaultFields = []FieldType{
		{
			Name: OrdinalField,
			Type: Int,
		},
		{
			Name: NameField,
			Type: String,
		},
	}
)

func (e *EnumType) toCode() []jen.Code {
	structName := MakeFirstLetterUpperCase(e.name)
	pluralStructName := MakePlural(structName)
	add := jen.
		Comment(fmt.Sprintf("%s type definition", structName)).Line().
		Type().Id(structName).Struct(e.fields.toCode()...).Line().
		Comment(pluralStructName).Line().
		Var().Defs(e.enumValues.toCode(structName)...).Line().
		Add(e.fields.GetMethods(structName)...).
		Add(getAllFunction(structName, e.enumValues)...)

	return *add
}

func (e *EnumType) addDefaultTypes() {
	e.fields = append(e.fields, defaultFields...)
}

func getDefaultValues(ordinal int, name string) []FieldValue {
	return []FieldValue{
		{
			Name:  OrdinalField,
			Value: ordinal,
		},
		{
			Name:  NameField,
			Value: name,
		},
	}
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

func (f FieldTypes) GetMethods(structName string) []jen.Code {
	statements := make([]jen.Code, 0)
	receiverName := GetFirstLetterInLowerCase(structName)

	for i := DefaultFieldCount; i < len(f); i++ {
		fieldType := f[i]
		id := jen.Func().Params(jen.Id(receiverName).Id(structName)).Id(MakeFirstLetterUpperCase(fieldType.Name)).
			Params().Id(fieldType.Type).Block(
			jen.Return(jen.Id(receiverName).Dot(fieldType.Name)),
		).Line().Line()

		statements = append(statements, id)
	}

	// add stringer methods
	stringerMethod := jen.Func().Params(jen.Id(receiverName).Id(structName)).Id("String").
		Params().Id("string").Block(
		jen.Return(jen.Qual("strconv", "Itoa").Call(jen.Id(receiverName).Dot(OrdinalField))),
	).Line().Line()
	statements = append(statements, stringerMethod)

	// add json Marshall method
	variableNameToBeEncoded := "val"
	marshallFunction := jen.Func().Params(jen.Id(receiverName).
		Id(structName)).Id("MarshalJSON").Params().Params(
		jen.Index().Byte(), jen.Error()).Block(
		jen.Id(variableNameToBeEncoded).Id(":=").Id(receiverName).Dot(OrdinalField).Line().Line().
			Return(jen.Qual("encoding/json", "Marshal").Params(jen.Id(variableNameToBeEncoded))),
	).Line().Line()
	statements = append(statements, marshallFunction)

	return statements
}

func (e EnumValues) toCode(structName string) []jen.Code {
	statements := make([]jen.Code, 0)

	for _, value := range e {
		block := jen.Line().
			Id(value.Name).Op("=").Id(structName).
			Block(value.fields.toCode()...).Line()

		statements = append(statements, block)
	}

	return statements
}

func (f FieldValues) toCode() []jen.Code {
	statements := make([]jen.Code, 0)

	for _, value := range f {
		code := jen.Id(MakeFirstLetterLowerCase(value.Name)).Id(":").Lit(value.Value).Id(",")
		statements = append(statements, code)
	}

	return statements
}
