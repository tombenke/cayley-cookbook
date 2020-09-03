package main

import (
	"context"
	"fmt"
	"github.com/cayleygraph/cayley/schema"
	"github.com/cayleygraph/quad"
	uuid "github.com/satori/go.uuid"
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
	ImportBookmarks()
}

func ImportBookmarks() {
	store := initStore()

	// Load Bookmarks from YAML and upload into the mem-store
	Import(store, bookmarksFixtures)

	// Check if Bookmarks have been uploaded
	stats, err := store.Stats(context.Background(), true)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("stats: %v\n", stats)
}
