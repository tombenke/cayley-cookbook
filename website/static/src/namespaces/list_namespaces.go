package main

import (
    "fmt"
	"github.com/cayleygraph/quad/voc"

    // Add some predefined vocabularies
	_ "github.com/cayleygraph/quad/voc/schema"
	_ "github.com/cayleygraph/quad/voc/xsd"
	_ "github.com/cayleygraph/quad/voc/rdf"
	_ "github.com/cayleygraph/quad/voc/rdfs"
)

func main() {
    fmt.Println("Registered namespaces:");
    for _, v := range voc.List() {
        fmt.Printf("  %s => %s\n", v.Prefix, v.Full);
    }
}
