package main

import (
	"context"
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
	// Create an in-memory store
	store := kbase.InitStore()

	// Import bookmarks data from yaml using the `Bookmark` schema
	ImportBookmarksWithSchema(store)

	// Find all Bookmarks, and convert to and array of object using `Bookmark` schema.
	bookmarks, _ := GetAllBookmarks(store)
	printBookmarks(bookmarks)

	// Find all Bookmarks, and convert to and array of object using `ShortBookmark` schema.
	shortBookmarks, _ := GetAllShortBookmarks(store)
	printShortBookmarks(shortBookmarks)
}

func ImportBookmarksWithSchema(store *cayley.Handle) {

	// Load Bookmarks from YAML and upload into the mem-store
	Import(store, bookmarksFixtures)

	// Check if Bookmarks have been uploaded
	_, err := store.Stats(context.Background(), true)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Successfully imported bookmarks into the mem-store.\n")
	}
}

func printBookmarks(bookmarks Bookmarks) {
	fmt.Println("\nFull Bookmarks:")
	for _, bookmark := range bookmarks {
		if bookmark.Author == "" {
			fmt.Printf("- \"%s\"\n", bookmark.Title)
		} else {
			fmt.Printf("- \"%s\" by %s\n\n", bookmark.Title, bookmark.Author)
		}
		fmt.Printf("  %s\n\n", bookmark.Urls[0])
	}
}

func printShortBookmarks(bookmarks Bookmarks) {
	fmt.Println("\nShort Bookmarks:")
	for _, bookmark := range bookmarks {
		fmt.Printf("- \"%s\"\n", bookmark.Title)
		fmt.Printf("  %s\n\n", bookmark.Urls[0])
	}
}
