# go-asset

## What is does

**go-asset** will embed a binary files as a variable containing a []byte litteral.

It is designed to be used from the command line, where the var name, the package name are specified from the command line with reasonable defaults.

The output file name is derived from the input file name. If it already exists, it is overwritten.

To get help, type :

    asset --help
