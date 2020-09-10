package main

import (
	"github.com/tombenke/cayley-cookbook-src/kbase/impex"
	"os"
)

func main() {

	inFile, outfile := getFileNames()

	quads, _ := impex.ImportFromFile(inFile, "")
	impex.ExportToFile(quads, outfile, "")
}

func getFileNames() (string, string) {
	args := os.Args[1:]
	if len(args) < 2 {
		panic("Too few arguments.")
	}
	return args[0], args[1]
}
