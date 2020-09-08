package main

import (
	"fmt"
	"github.com/cayleygraph/quad"
	"reflect"

	// Add some predefined vocabularies
	"github.com/cayleygraph/quad/voc/rdf"
	"github.com/tombenke/cayley-cookbook-src/kbase/voc/foaf"
)

func main() {

	// Create Quads about some people using Blank Nodes as references
	quadsWithBNodes := makeQuadsWithBNodes()

	// Take a look at the resulted quads
	printQuads(quadsWithBNodes)

	// Create Quads about some people using IRIs as references
	quadsWithIRIs := makeQuadsWithIRIs()

	// Take a look at the resulted quads
	printQuads(quadsWithIRIs)
}

func makeQuadsWithBNodes() []quad.Quad {
	label := "people"
	quads := []quad.Quad{}

	// Generate a Blank Node to represent the person internally
	// You can create a sequence of internal IDs for the Blank Nodes
	var seq quad.Sequence
	luke := seq.Next()

	// Alternatively create a Blank Node with a random internal ID
	// luke := quad.RandomBlankNode()

	quads = append(quads, quad.Make(luke, quad.IRI(rdf.Type), quad.IRI(foaf.Person), label))
	quads = append(quads, quad.Make(luke, quad.IRI(foaf.FamilyName), "Luke", label))
	quads = append(quads, quad.Make(luke, quad.IRI(foaf.GivenName), "Skywalker", label))
	quads = append(quads, quad.Make(luke, quad.IRI(foaf.Age), 23, label))

	leia := seq.Next()

	quads = append(quads, quad.Make(leia, quad.IRI(rdf.Type), quad.IRI(foaf.Person), label))
	quads = append(quads, quad.Make(leia, quad.IRI(foaf.FamilyName), "Leia", label))
	quads = append(quads, quad.Make(leia, quad.IRI(foaf.GivenName), "Organa", label))

	quads = append(quads, quad.Make(leia, quad.IRI(foaf.Knows), luke, label))

	return quads
}

func makeQuadsWithIRIs() []quad.Quad {
	label := "people"
	quads := []quad.Quad{}

	// Create IRIs to represent the person globally, and universally
	luke := quad.IRI("https://swapi.co/resource/human/luke_skywalker")

	// Alternatively create a Blank Node with a random internal ID
	// luke := quad.RandomBlankNode()

	quads = append(quads, quad.Make(luke, quad.IRI(rdf.Type), quad.IRI(foaf.Person), label))
	quads = append(quads, quad.Make(luke, quad.IRI(foaf.FamilyName), "Luke", label))
	quads = append(quads, quad.Make(luke, quad.IRI(foaf.GivenName), "Skywalker", label))

	leia := quad.IRI("https://swapi.co/resource/human/leia_organa")

	quads = append(quads, quad.Make(leia, quad.IRI(rdf.Type), quad.IRI(foaf.Person), label))
	quads = append(quads, quad.Make(leia, quad.IRI(foaf.FamilyName), "Leia", label))
	quads = append(quads, quad.Make(leia, quad.IRI(foaf.GivenName), "Organa", label))

	quads = append(quads, quad.Make(leia, quad.IRI(foaf.Knows), luke, label))

	return quads
}

func printQuads(quads []quad.Quad) {
	fmt.Println("The details of the quads created:")
	for i, q := range quads {
		fmt.Printf("quads[%d]:\n", i)
		fmt.Printf("  subject: %s %v\n", q.Get(quad.Subject), reflect.TypeOf(q.Get(quad.Subject)))
		fmt.Printf("  predicate: %s %v\n", q.Get(quad.Predicate), reflect.TypeOf(q.Get(quad.Predicate)))
		fmt.Printf("  object: %s %v\n", q.Get(quad.Object), reflect.TypeOf(q.Get(quad.Object)))
		fmt.Printf("  label: %s %v\n\n", q.Get(quad.Label), reflect.TypeOf(q.Get(quad.Label)))
	}

	fmt.Println("The quads in NQuad representation:")
	for _, q := range quads {
		fmt.Printf("%s\n", q.NQuad())
	}
}
