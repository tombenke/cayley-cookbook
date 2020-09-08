package main

import (
    "fmt"
    "log"
    "reflect"
    "context"

    "github.com/cayleygraph/cayley"
)

func main() {

    // Create a brand new graph
    store, err := cayley.NewMemoryGraph()
    if err != nil {
        log.Fatalln(err)
    }

    // Create Quads about some people using Blank Nodes as references
    quads := makeQuads()

    // Add quads to the store
    for _, q := range(quads) {
        store.AddQuad(q)
    }

    printAllNodes(store)
}

func printAllNodes(store *cayley.Handle) {
    // Get the iterator that enumerates all nodes in the graph.
    it := store.NodesAllIterator()
    fmt.Printf("%v %v\n", it, reflect.TypeOf(it))
    for it.Next(context.Background()) {
        ref := it.Result()
        key := ref.Key()
        value := store.NameOf(ref)
        fmt.Printf("Ref: %v %v\n", ref, reflect.TypeOf(ref))
        fmt.Printf("Key: %v %v\n", key, reflect.TypeOf(key))
        fmt.Printf("Value: %v %v\n", value, reflect.TypeOf(value))
        fmt.Println()
    }
}
