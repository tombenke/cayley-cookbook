package main

import (
	"fmt"
	"github.com/tombenke/cayley-cookbook-src/kbase/impex"
)

func main() {
	quads, _ := impex.ImportFromFile("../../data/testdata.nq", "nquads")
	fmt.Printf("%v\n", quads)
}
