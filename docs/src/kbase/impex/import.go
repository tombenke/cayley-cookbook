package impex

import (
	"github.com/cayleygraph/quad"
	_ "github.com/cayleygraph/quad/nquads"
	"io"
	"os"
)

func ImportFromFile(inFileName string, formatName string) ([]quad.Quad, error) {

	// Get formatter
	format := GetFormatter(inFileName, formatName)

	// Open the input file.
	// If the `inFileName` is "" then open then returns with `os.Stdin`.
	var fileReader io.Reader

	if inFileName == "" {
		fileReader = os.Stdin
	} else {
		file, err := os.Open(inFileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		fileReader = file
	}

	// Add the fileWriter to the formatter
	quadReader := format.Reader(fileReader)
	defer quadReader.Close()

	// Make a `QuadWriter` to `quads` array
	var quads []quad.Quad
	quadWriter := NewQuadWriter(&quads)

	// Execute the conversion and writing of quads
	_, err := quad.Copy(quadWriter, quadReader)

	if err != nil {
		panic(err)
	}

	return quads, nil
}

type QuadWriter struct {
	quads *[]quad.Quad
}

func (qw QuadWriter) WriteQuad(buf quad.Quad) error {
	*qw.quads = append((*qw.quads), buf)
	return nil
}

func (qw QuadWriter) WriteQuads(buf []quad.Quad) (int, error) {
	for _, q := range buf {
		*qw.quads = append((*qw.quads), q)
	}
	return len(buf), nil
}

func NewQuadWriter(quads *[]quad.Quad) quad.Writer {
	qw := new(QuadWriter)
	qw.quads = quads

	return qw
}
