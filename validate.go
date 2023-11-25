package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

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
