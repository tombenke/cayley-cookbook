package main

import (
	"github.com/cayleygraph/quad"

	// Add some predefined vocabularies
	"github.com/cayleygraph/quad/voc/rdf"
	"github.com/tombenke/cayley-cookbook-src/kbase/voc/foaf"
)

func makeQuads() []quad.Quad {
	// Generate Blank Nodes to represent the people internally
	luke := quad.RandomBlankNode()
	leia := quad.IRI("starwars:leia_organa") //quad.RandomBlankNode()

	// Create Quads about to export
	label := "people"
	quads := []quad.Quad{}
	quads = append(quads, quad.Make(luke, quad.IRI(rdf.Type), quad.IRI(foaf.Person), label))
	quads = append(quads, quad.Make(luke, quad.IRI(foaf.GivenName), "Luke", label))
	quads = append(quads, quad.Make(luke, quad.IRI(foaf.FamilyName), "Skywalker", label))
	quads = append(quads, quad.Make(luke, quad.IRI(foaf.Age), 23, label))

	quads = append(quads, quad.Make(leia, quad.IRI(rdf.Type), quad.IRI(foaf.Person), label))
	quads = append(quads, quad.Make(leia, quad.IRI(foaf.Knows), luke, label))
	quads = append(quads, quad.Make(leia, quad.IRI(foaf.GivenName), "Leia", label))
	quads = append(quads, quad.Make(leia, quad.IRI(foaf.FamilyName), "Organa", label))

	return quads
}
