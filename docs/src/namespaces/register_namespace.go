package main

import (
	"fmt"
	"github.com/cayleygraph/quad/voc"

	// Add built-in vocabulary
	_ "github.com/cayleygraph/quad/voc/rdfs"

	// Add own vocabulary
	_ "github.com/tombenke/cayley-cookbook-src/kbase/voc/foaf"
)

func init() {
	// Register a new namespace with prefix
	voc.RegisterPrefix(`acc:`, `http://mycompany.com/voc/accounting#`)
}

func main() {
	fmt.Println("Registered namespaces:")
	for _, v := range voc.List() {
		fmt.Printf("  %s => %s\n", v.Prefix, v.Full)
	}
}
