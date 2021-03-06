//Package main : create a source go file, that encapsulate any binary data.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Suffix added to the file name to get the output file name
var suffix string

// VarName defines the name of the []byte variable
var varName string

// PackName is the name of the package of the output file
var packName string

var help bool

func init() {
	flag.StringVar(&varName, "var", "mybuffer", "the name of the variable containing the []bytes. Can be local or global.")
	flag.StringVar(&varName, "v", "mybuffer", "the name of the variable containing the []bytes. Can be local or global.")
	flag.StringVar(&packName, "package", "myassets", "the package name where the data will be available")
	flag.StringVar(&packName, "p", "myassets", "the package name where the data will be available")
	flag.StringVar(&suffix, "suffix", "_ast.go", "the suffix added to the input file name. Should end with .go to be useful ...")
	flag.StringVar(&suffix, "s", "_ast.go", "the suffix added to the input file name. Should end with .go to be useful ...")
	flag.BoolVar(&help, "help", false, "display this help information")
	flag.BoolVar(&help, "h", false, "display this help information")
}

func main() {
	flag.Parse()
	if help {
		printHelp()
		return
	}
	if flag.NArg() != 1 || len(flag.Args()[0]) == 0 {
		fmt.Println("Error : exactly one input file expected")
		printHelp()
		return
	}
	createAsset()
}

func printHelp() {
	fmt.Printf("Usage : %s [options ...] inputFileName\n", os.Args[0])
	fmt.Println("Version 1.0 - (c) 2019 - Xavier Gandillot")
	flag.PrintDefaults()
}

// CreateAsset creates a new go asset file
func createAsset() {
	ifn := flag.Args()[0]
	fmt.Printf("Processing %s --> %s\n", ifn, ifn+suffix)
	b, err := ioutil.ReadFile(ifn)
	if err != nil {
		panic(err)
	}
	if len(suffix) == 0 {
		panic("Attempting to override the provided file - Suffix cannot be empty !")
	}
	f, err := os.Create(ifn + suffix)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintf(f,
		"//Package %s contains autogenerated data\n//Do not edit this file - automatically generated\n//Source file : %s\n//Generated on %v",
		packName,
		ifn,
		time.Now())
	fmt.Fprintf(f, "\npackage %s\n//%s is the []byte slice\nvar %s []byte\nfunc init() {",
		packName,
		varName,
		varName)
	fmt.Fprintf(f, "\n%s = []byte{", varName)

	for i, a := range b {
		if i%20 == 0 {
			fmt.Fprintf(f, "\n")
		}
		fmt.Fprintf(f, "%d,", a)
	}
	fmt.Fprintf(f, "}}\n")
}
