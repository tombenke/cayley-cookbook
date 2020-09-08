package main

import (
	"github.com/tombenke/cayley-cookbook-src/kbase/impex"
)

func main() {

	// Create some data to export
	quads := makeData()

	// Export in `nquads` format
	// to the `stdout`
	impex.ExportToFile(quads, "", "nquads")

	// to a file with explicit format definition by name
	impex.ExportToFile(quads, "./people.nq", "nquads")

	// to a file with implicit format definition via the file extension
	impex.ExportToFile(quads, "./people.nq", "")

	// to a file with implicit format definition via mime-type
	impex.ExportToFile(quads, "./people.nq", "application/n-triples")

	// Export in `JSON-LD` format
	impex.ExportToFile(quads, "./people.jsonld", "")

	// Export in `JSON` format
	impex.ExportToFile(quads, "./people.json", "")

	// Export in `GraphML` format
	impex.ExportToFile(quads, "./people.graphml", "")

	// Export in `GML` format
	impex.ExportToFile(quads, "./people.gml", "")

	// Export in `Graphviz` format
	impex.ExportToFile(quads, "./people.dot", "")
}
