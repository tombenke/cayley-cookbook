package main

import (
	"github.com/tombenke/cayley-cookbook-src/kbase/impex"
	"os"
)

func main() {

	inFile, outfile := getFileNames()

	quads, err := impex.ReadFromTurtle(inFile, "", false)
	if err != nil {
		panic(err)
	}

	impex.ExportToFile(quads, outfile, "")
}

func getFileNames() (string, string) {
	args := os.Args[1:]
	if len(args) < 2 {
		panic("Too few arguments.")
	}
	return args[0], args[1]
}
