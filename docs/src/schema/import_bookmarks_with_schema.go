package main

import (
	"context"
	"fmt"
	"github.com/cayleygraph/cayley/schema"
	"github.com/cayleygraph/quad"
	uuid "github.com/satori/go.uuid"
	"github.com/tombenke/cayley-cookbook-src/kbase"
)

const bookmarksFixtures = "./bookmarks.yml"

func init() {
	schema.GenerateID = func(o interface{}) quad.Value {
		switch o.(type) {
		case Bookmark:
			return quad.IRI(o.(Bookmark).ID)
		default:
			return quad.IRI(uuid.NewV4().String())
		}
	}
}

func main() {
	ImportBookmarksWithSchema()
}

func ImportBookmarksWithSchema() {
	// Create an in-memory store
	store := kbase.InitStore()

	// Load Bookmarks from YAML and upload into the mem-store
	Import(store, bookmarksFixtures)

	// Check if Bookmarks have been uploaded
	stats, err := store.Stats(context.Background(), true)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Successfully imported bookmarks into the mem-store as %d nodes and %d quads.\n", stats.Nodes.Size, stats.Quads.Size)
	}
}
