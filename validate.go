package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const GoFileExtension = ".go"

func validate(source, destination, pkg *string) error {
	if err := validateSource(source); err != nil {
		return err
	}
	if err := validateFileName(destination); err != nil {
		return err
	}
	if pkg == nil {
		return errors.New("package cannot be nil")
	}
	*pkg = strings.TrimSpace(*pkg)
	if len(*pkg) == 0 {
		return errors.New("pkg must be set")
	}

	return nil
}

func validateSource(source *string) error {
	if err := validateFileName(source); err != nil {
		return fmt.Errorf("source is not valid: %w", err)
	}
	if _, err := os.Open(*source); err != nil {
		return fmt.Errorf("could not open the file: %w", err)
	}

	return nil
}

func validateFileName(fileName *string) error {
	if fileName == nil {
		return errors.New("variable is nil")
	}
	*fileName = strings.TrimSpace(*fileName)
	if len(*fileName) == 0 {
		return errors.New("file name cannot be empty")
	}
	if !strings.HasSuffix(*fileName, GoFileExtension) {
		return errors.New("file must be go file")
	}

	return nil
}
