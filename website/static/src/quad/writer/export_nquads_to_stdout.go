package main

import (
    "fmt"
    "os"
    "io"
    "strings"
	"path/filepath"
	"github.com/cayleygraph/quad"

    // Load formatters
	_ "github.com/cayleygraph/quad/dot"
	_ "github.com/cayleygraph/quad/gml"
//	_ "github.com/cayleygraph/quad/graphml"
	_ "github.com/cayleygraph/quad/json"
	_ "github.com/cayleygraph/quad/jsonld"
	_ "github.com/cayleygraph/quad/nquads"
)

func main() {

    // Create some data to export
    quads := makeData()

    // Export in `nquads` format
    // to the `stdout`
    exportToFile(quads, "", "nquads")

    // to a file with explicit format definition by name
    exportToFile(quads, "./people.nq", "nquads")

    // to a file with implicit format definition via the file extension
    exportToFile(quads, "./people.nq", "")

    // to a file with implicit format definition via mime-type
    exportToFile(quads, "./people.nq", "application/n-triples")

    // Export in `JSON-LD` format
    exportToFile(quads, "./people.jsonld", "")

    // Export in `JSON` format
    exportToFile(quads, "./people.json", "")

    // Export in `GraphML` format
    exportToFile(quads, "./people.graphml", "")

    // Export in `GML` format
    exportToFile(quads, "./people.gml", "")

    // Export in `Graphviz` format
    exportToFile(quads, "./people.dot", "")
}

// Export `quads` into a file named `outFileName` in `formatName` representational format.
// If `outFileName` equals to "" then the results will be written to the `os.Stdout`.
// The following formats can be identified via several methods:
//      |-----------|-----------|------------------------|
//      | name      | extension |   mime-type            |
//      |-----------|-----------|------------------------|
//      | nquads    | .nq, .nt  | application/n-quads,   |
//      |           |           | application/n-triplets |
//      |-----------|-----------|------------------------|
//      | jsonld    | .jsonld   | application/ld+json    |
//      |-----------|-----------|------------------------|
//      | json      | .json     | application/json       |
//      |-----------|-----------|------------------------|
//      | graphml   | .graphml  | application/xml        |
//      |-----------|-----------|------------------------|
//      | gml       | .gml      |                        |
//      |-----------|-----------|------------------------|
//      | graphviz  | .gv, .dot |                        |
//      |-----------|-----------|------------------------|
//   
func exportToFile(quads []quad.Quad, outFileName string, formatName string) {

    // Get formatter
    format := getFormatter(outFileName, formatName)
    
    // Open/create the output file.
    // If the `outFileName` is "" then open then returns with `os.Stdout`.
	var fileWriter io.Writer

	if outFileName == "" {
		fileWriter = os.Stdout
	} else {
		file, err := os.Create(outFileName)
		if err != nil {
			panic(err)
		}
        defer file.Close()
		fileWriter = file
	}

    // Add the fileWriter to the formatter
    quadWriter := format.Writer(fileWriter)
	defer quadWriter.Close()

    // Make a `QuadReader` from the `quads` array
    quadReader := quad.NewReader(quads)

    // Execute the conversion and writing of quads
	_, err := quad.Copy(quadWriter, quadReader)

	if err != nil {
		panic(err)
	} else if err = quadWriter.Close(); err != nil {
		panic(err)
	}
}

// Get the formatter selected either by the `formatName` or by the file extension.
func getFormatter(outFileName string, formatName string) (format *quad.Format) {
	if formatName != "" {
        if strings.Contains(formatName, "/") {
		    format = quad.FormatByMime(formatName)
        } else {
		    format = quad.FormatByName(formatName)
        }
	} else {
        format = formatByFileName(outFileName)
    }

	if format == nil {
		panic(fmt.Sprintf("unsupported format: %q", formatName))
	} else if format.Writer == nil {
		panic(fmt.Sprintf("encoding in %s format is not supported", formatName))
	}

    return format
}

// Determine the representational format from the file extension
func formatByFileName(fileName string) *quad.Format {
	ext := filepath.Ext(fileName)
	return quad.FormatByExt(ext)
}
