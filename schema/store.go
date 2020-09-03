package main

import (
	"context"
	"fmt"
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	_ "github.com/cayleygraph/cayley/graph/kv/bolt"
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

func initPStore() *cayley.Handle {
	var dbFileName = "db.boltdb"
	graph.InitQuadStore("bolt", dbFileName, nil)

	// Open and use the database
	store, err := cayley.NewGraph("bolt", dbFileName, nil)
	if err != nil {
		log.Fatalln("Error in NewGraph: ", err)
	}

	return store
}

func printQuads(store *cayley.Handle) {
	// get all quads
	it := store.QuadsAllIterator()
	defer it.Close()

	fmt.Println("Quads:")
	fmt.Println("-----")

	ctx := context.TODO()

	for it.Next(ctx) {
		fmt.Println(store.Quad(it.Result()))
	}

	fmt.Println()
}
