//Package asset allows to create a source go file, that encapsulate any binary data.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Suffix added to the file name to get the output file name
var Suffix string

// VarName defines the name of the []byte variable
var VarName string

// PackName is the name of the package of the output file
var PackName string

var help bool

func init() {
	flag.StringVar(&VarName, "var", "BinaryDataBuffer", "the name of the variable containig the []bytes. Can be local or global.")
	flag.StringVar(&PackName, "package", "assets", "the package name where the data will be available")
	flag.StringVar(&Suffix, "suffix", "_asset.go", "the suffix added to the input file name. Should end with .go to be useful ...")
	flag.BoolVar(&help, "help", false, "display this help information")
}

func main() {
	flag.Parse()
	//fmt.Println(VarName, PackName, Suffix, help, InputFileName, flag.Args())
	if flag.NArg() != 1 {
		panic("Exactly one input file expected")
	}
	if help {
		fmt.Println("go-asset [options ...] inputFileName")
		flag.PrintDefaults()
		return
	}

}

// CreateAsset creates a new go asset file
func CreateAsset() {
	ifn := flag.Args()[0]
	fmt.Printf("\nProcessing %s --> %s", ifn, ifn+Suffix)
	b, err := ioutil.ReadFile(ifn)
	if err != nil {
		panic(err)
	}
	if len(Suffix) == 0 {
		panic("Attempting to override the provided file - Suffix cannot be empty !")
	}
	f, err := os.Create(ifn + Suffix)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintf(f,
		"//Package %s\n// Automatically generated - do not edit\n//Source file : %s\n//Date : %v",
		PackName,
		ifn,
		time.Now())
	fmt.Fprintf(f, "\npackage asset\n//%s is the []byte slice\nvar %s []byte\nfunc init() {",
		VarName,
		VarName)
	fmt.Fprintf(f, "\n%s = []byte{", VarName)

	for i, a := range b {
		if i%20 == 0 {
			fmt.Fprintf(f, "\n")
		}
		fmt.Fprintf(f, "%d,", a)
	}
	fmt.Fprintf(f, "}}\n")
}
