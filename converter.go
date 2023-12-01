package main

import (
	"fmt"
	"strconv"
)

type TypeConverter func(string) (any, error)

var typesAndConverters = make(map[string]TypeConverter)

const (
	Int     = "int"
	Int8    = "int8"
	Int16   = "int16"
	Int32   = "int32"
	Int64   = "int64"
	Uint    = "uint"
	Uint8   = "uint8"
	Uint16  = "uint16"
	Uint32  = "uint32"
	Uint64  = "uint64"
	Float32 = "float32"
	Float64 = "float64"
	String  = "string"
	Bool    = "bool"
)

func init() {
	typesAndConverters[String] = stringConverter
	typesAndConverters[Int] = intConverter
	typesAndConverters[Int8] = intConverter
	typesAndConverters[Int16] = intConverter
	typesAndConverters[Int32] = intConverter
	typesAndConverters[Int64] = intConverter
	typesAndConverters[Uint] = intConverter
	typesAndConverters[Uint8] = intConverter
	typesAndConverters[Uint16] = intConverter
	typesAndConverters[Uint32] = intConverter
	typesAndConverters[Uint64] = intConverter
	typesAndConverters[Float32] = floatConverter
	typesAndConverters[Float64] = floatConverter
	typesAndConverters[Bool] = boolConverter
}

func stringConverter(s string) (any, error) {
	return s, nil
}

func intConverter(s string) (any, error) {
	return strconv.Atoi(s)
}

func floatConverter(s string) (any, error) {
	return strconv.ParseFloat(s, 64)
}

func boolConverter(s string) (any, error) {
	return strconv.ParseBool(s)
}

func convert(valueType, value string) (any, error) {
	converter, present := typesAndConverters[valueType]
	if !present {
		return nil, fmt.Errorf("%s is not a supported type", valueType)
	}

	parsedValue, err := converter(value)
	if err != nil {
		return nil, fmt.Errorf("could not parse %s as %s :%w", value, valueType, err)
	}

	return parsedValue, nil
}
