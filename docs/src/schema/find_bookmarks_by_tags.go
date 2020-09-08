package main

import (
	"fmt"
	"github.com/cayleygraph/cayley"
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
	store := ImportBookmarksWithSchema()

	bookmarks, _ := FindBookmarksByAnyTag(store, []quad.Value{quad.String("nosql"), quad.String("gremlin")})
	fmt.Println("\nResults of finding any of `nosql` or `gremlin` tags:")
	printResults(bookmarks)

	bookmarks, _ = FindBookmarksByEveryTag(store, []quad.Value{quad.String("cayley"), quad.String("gremlin")})
	fmt.Println("\nResults of finding every tags of `cayley` or `gremlin`:")
	printResults(bookmarks)
}

func ImportBookmarksWithSchema() *cayley.Handle {
	// Create an in-memory store
	store := kbase.InitStore()

	// Load Bookmarks from YAML and upload into the mem-store
	Import(store, bookmarksFixtures)

	return store
}

func printResults(bookmarks Bookmarks) {
	for _, bookmark := range bookmarks {
		fmt.Printf("- \"%s\"\n", bookmark.Title)
		fmt.Printf("  %v\n", bookmark.Tags)
		fmt.Printf("  %s\n\n", bookmark.Urls[0])
	}
}
