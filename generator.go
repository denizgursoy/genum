package main

import (
	"os"
	"path/filepath"

	"github.com/dave/jennifer/jen"
)

func generate(t []EnumType, pkg string) string {
	f := jen.NewFilePath(pkg)
	f.HeaderComment(GeneratedMessage)
	for _, enumType := range t {
		f.Add(enumType.toCode()...)
	}

	return f.GoString()
}

func write(targetFile, content string) error {
	if err := os.MkdirAll(filepath.Dir(targetFile), os.ModePerm); err != nil {
		return err
	}
	createdFile, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	if _, err := createdFile.WriteString(content); err != nil {
		return err
	}

	return nil
}
