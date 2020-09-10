package main

import (
	"log"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/quad"
	"github.com/tombenke/cayley-cookbook-src/kbase"
	"github.com/tombenke/cayley-cookbook-src/kbase/impex"
)

func main() {
	store := kbase.InitStore()

	quads, err := impex.ReadFromTurtle("../../data/starwars.ttl", "http://example.org", false)
	if err != nil {
		log.Fatal(err)
	}

	for _, q := range quads {
		store.AddQuad(q)
	}

	numberOfPlanets := countPlanets(store)

	log.Println("Number of known planets:", numberOfPlanets)
}

func countPlanets(store *cayley.Handle) int64 {
	var a int64
	p := cayley.StartPath(store).HasReverse(quad.String("<https://swapi.co/vocabulary/planet>")).Unique().Count()

	p.Iterate(nil).EachValue(nil, func(value quad.Value) {
		a = value.Native().(int64)
	})

	return a
}
