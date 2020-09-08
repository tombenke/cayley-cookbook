package main

import (
    "fmt"
    "log"
    "reflect"

    "github.com/cayleygraph/cayley"
    "github.com/cayleygraph/quad"
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

    queryAllNodes(store)
}

func queryAllNodes(store *cayley.Handle) {

    // Now we create the path, to get to our data
    p := cayley.StartPath(store)
    fmt.Printf("p.TypeOf: %v\n", reflect.TypeOf(p))
    fmt.Printf("p.Iterate(nil).TypeOf: %v\n", reflect.TypeOf(p.Iterate(nil)))

    // Now we iterate over results.
    err := p.Iterate(nil).EachValue(nil, func(value quad.Value){
        fmt.Printf("%v %v\n", value, reflect.TypeOf(value))
    })
    if err != nil {
        log.Fatalln(err)
    }
}
