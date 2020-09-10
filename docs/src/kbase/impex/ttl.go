package impex

import (
	"log"
	"os"
	path "path/filepath"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/quad"
	rdf "github.com/deiu/gon3"
)

// ReadFromTTL reads a ttl file and return a slice of quads
func ReadFromTurtle(ttlPath string, baseURI string, useShortIRI bool) ([]quad.Quad, error) {

	absPath, err := path.Abs(ttlPath)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	inputFile, err := os.Open(absPath)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	
	parser := rdf.NewParser(baseURI)

	graph, err := parser.Parse(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	quads := []quad.Quad{}

	triples := graph.IterTriples()

	for t := range triples {
		if useShortIRI {
			s := quad.IRI(t.Subject.RawValue())
			p := quad.IRI(t.Predicate.RawValue())
			o := quad.IRI(t.Object.RawValue())

			q := cayley.Triple(quad.IRI.Short(s), quad.IRI.Short(p), quad.IRI.Short(o))
			quads = append(quads, q)
		} else {
			quads = append(quads, cayley.Triple(t.Subject, t.Predicate, t.Object))
		}
	}

	return quads, nil
}
