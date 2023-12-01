package main

import "github.com/fatih/color"

var (
	errorPrinter   = color.New(color.FgRed)
	successPrinter = color.New(color.FgCyan)
)

func PrintError(msg string) {
	errorPrinter.Println(msg)
}

func PrintSuccess(msg string) {
	successPrinter.Println(msg)
}
