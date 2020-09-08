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

    printAllQuads(store)
}

func printAllQuads(store *cayley.Handle) {
    // Get the iterator that enumerates all nodes in the graph.
    it := store.QuadsAllIterator()
    //fmt.Printf("%v %v\n", it, reflect.TypeOf(it))
    for it.Next(context.Background()) {
        ref := it.Result()
        key := ref.Key()
        value := store.NameOf(ref)
        quad := store.Quad(ref)
        fmt.Printf("Ref: %v %v\n", ref, reflect.TypeOf(ref))
        fmt.Printf("Key: %v %v\n", key, reflect.TypeOf(key))
        fmt.Printf("Value: %v %v\n", value, reflect.TypeOf(value))
        fmt.Printf("Quad: %v %v\n", quad, reflect.TypeOf(quad))
        fmt.Println()
        /*
        v := it.Result()
        fmt.Printf("%v %v\n", v, reflect.TypeOf(v))
        fmt.Printf("%v\n", v.Key())

        name := store.NameOf(v)
        fmt.Printf("%v %v\n", name, reflect.TypeOf(name))
        q := store.Quad(v)
        fmt.Println(q)
        fmt.Println()
        */
    }
}
