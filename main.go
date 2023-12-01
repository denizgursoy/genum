package main

import (
	"flag"
	"fmt"
)

func main() {
	source := flag.String("source", "", "path of source file")
	destination := flag.String("destination", "enum.go", "path of destination file")
	pkg := flag.String("package", "enums", "package in destination file")
	flag.Parse()

	if err := validate(source, destination, pkg); err != nil {
		fmt.Println(err.Error())
		return
	}

	types, err := parseSource(*source)
	if err != nil {
		PrintError(err.Error())
		return
	}

	for _, enumType := range types {
		PrintSuccess(fmt.Sprintf("%s type is found", enumType.name))
	}

	content := generate(types, *pkg)

	if err := write(*destination, content); err != nil {
		PrintError(err.Error())
		return
	}

	PrintSuccess(fmt.Sprintf("%s is created", *destination))
}
