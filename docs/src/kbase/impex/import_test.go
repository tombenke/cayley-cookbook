package impex

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadNquadsFromFile(t *testing.T) {
	quads, err := ImportFromFile("../../data/testdata.nq", "nquads")
	fmt.Printf("%v\n", quads)
	assert.Equal(t, err, nil)
}
