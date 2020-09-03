package main

import (
	"github.com/cayleygraph/cayley"
	"log"
)

// Create a mem-store for testing
func initStore() *cayley.Handle {
	store, err := cayley.NewMemoryGraph()
	if err != nil {
		log.Fatalln(err)
	}

	return store
}
