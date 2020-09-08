package main

import (
	"context"
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/schema"
	"github.com/cayleygraph/quad"
	"github.com/tombenke/cayley-cookbook-src/kbase/impex"
	"log"
)

type Bookmarks []Bookmark

type Bookmark struct {
	ID        quad.IRI `yaml:"id" quad:"id"`
	Title     string   `yaml:"title" quad:"title"`
	Content   string   `yaml:"content" quad:"content,optional"`
	Author    string   `yaml:"author" quad:"author,optional"`
	Published string   `yaml:"published" quad:"published,optional"`
	Tags      []string `yaml:"tags,flow" quad:"tags"`
	Urls      []string `yaml:"urls,flow" quad:"urls"`
}

type BookmarkShort struct {
	ID    quad.IRI `yaml:"id" quad:"id"`
	Title string   `yaml:"title" quad:"title"`
	Urls  []string `yaml:"urls,flow" quad:"urls"`
}

func init() {
	schema.RegisterType("Bookmark", Bookmark{})
	schema.RegisterType("BookmarkShort", BookmarkShort{})
}

func Import(store *cayley.Handle, dbPath string) {
	var bm Bookmarks
	impex.ReadFromYaml(dbPath, &bm)
	qw := graph.NewWriter(store)
	defer qw.Close() // don't forget to close a writer; it has some internal buffering
	for _, b := range bm {
		_, err := schema.WriteAsQuads(qw, b)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}
}

func GetAllShortBookmarks(store *cayley.Handle) (Bookmarks, error) {
	var bms Bookmarks
	//p := cayley.StartPath(store).Has(quad.IRI("rdf:type"), quad.IRI("Bookmark"))
	p := cayley.StartPath(store, quad.IRI("Bookmark")).In(quad.IRI("rdf:type"))
	schemaConfig := schema.NewConfig()
	err := p.Iterate(context.Background()).EachValuePair(nil, func(ref graph.Ref, value quad.Value) {
		var b Bookmark
		schemaConfig.LoadTo(context.Background(), store, &b, value)
		//log.Printf("%v, %v, %v\n", reflect.TypeOf(ref), ref.Key(), b)
		bms = append(bms, b)
	})

	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	return bms, nil
}

func GetAllBookmarks(store *cayley.Handle) (Bookmarks, error) {
	var bms Bookmarks
	p := cayley.StartPath(store, quad.IRI("Bookmark")).In(quad.IRI("rdf:type"))
	schemaConfig := schema.NewConfig()
	err := p.Iterate(context.Background()).EachValuePair(nil, func(ref graph.Ref, value quad.Value) {
		var b Bookmark
		schemaConfig.LoadTo(context.Background(), store, &b, value)
		bms = append(bms, b)
	})

	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	return bms, nil
}

func FindBookmarksByAnyTag(store *cayley.Handle, tags []quad.Value) (Bookmarks, error) {
	schemaConfig := schema.NewConfig()
	var bms Bookmarks
	p := cayley.StartPath(store).Has(quad.IRI("rdf:type"), quad.IRI("Bookmark")).Has(quad.IRI("tags"), tags...)
	err := p.Iterate(context.Background()).EachValue(nil, func(value quad.Value) {
		var b Bookmark
		schemaConfig.LoadTo(context.Background(), store, &b, value)
		bms = append(bms, b)
	})
	return bms, err
}

func FindBookmarksByEveryTag(store *cayley.Handle, tags []quad.Value) (Bookmarks, error) {
	schemaConfig := schema.NewConfig()
	var bms Bookmarks
	bookmarksMorph := cayley.StartPath(store).Has(quad.IRI("rdf:type"), quad.IRI("Bookmark"))
	p := bookmarksMorph
	for _, t := range tags {
		p = cayley.StartPath(store).Follow(p).Has(quad.IRI("tags"), t)
	}
	err := p.Iterate(context.Background()).EachValue(nil, func(value quad.Value) {
		var b Bookmark
		schemaConfig.LoadTo(context.Background(), store, &b, value)
		bms = append(bms, b)
	})
	return bms, err
}
