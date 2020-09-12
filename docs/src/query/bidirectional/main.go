package main

import (
	"context"
	"fmt"
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/quad"
)

func main() {
	// Create an in-memory store
	store := InitStore()
	// Create Quads and uploads to the store
	quads := makeQuads()
	AddQuadsToStore(store, quads)

	// Execute the query
	doQuery(store)
}

// Create a mem-store for testing
func InitStore() *cayley.Handle {
	store, err := cayley.NewMemoryGraph()
	if err != nil {
		panic(err)
	}

	return store
}

func makeQuads() []quad.Quad {
	quads := []quad.Quad{}

	quads = append(quads, quad.Make("A", "1", "B", ""))
	quads = append(quads, quad.Make("B", "1", "A", ""))
	quads = append(quads, quad.Make("B", "2", "C", ""))
	quads = append(quads, quad.Make("C", "2", "B", ""))

	quads = append(quads, quad.Make("A1", "1", "B1", ""))
	quads = append(quads, quad.Make("B1", "1", "A1", ""))
	quads = append(quads, quad.Make("B1", "2", "C1", ""))
	quads = append(quads, quad.Make("C1", "2", "B1", ""))

	quads = append(quads, quad.Make("A2", "1", "B2", ""))
	quads = append(quads, quad.Make("B2", "1", "A2", ""))
	quads = append(quads, quad.Make("B2", "2", "C2", ""))
	quads = append(quads, quad.Make("C2", "2", "B2", ""))

	return quads
}

func AddQuadsToStore(store *cayley.Handle, quads []quad.Quad) {
	for _, q := range quads {
		store.AddQuad(q)
	}
}

func doQuery(store *cayley.Handle) {
	p := cayley.StartPath(store).Tag("source").Out("1").Out("2").Tag("target").Out("2").Out("1")

	p.Iterate(context.Background()).EachValue(nil, func(value quad.Value) {
		fmt.Printf("%v\n", value)
	})
	p.Iterate(context.Background()).TagValues(store, func(t map[string]quad.Value) {
		fmt.Printf("%v\n", t)
	})
}
