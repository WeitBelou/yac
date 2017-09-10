package yac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Returns handler without any changes
var emptyWrapper = WrapperFunc(func(h Handler) Handler {
	return h
})

func TestWrapper_Add(t *testing.T) {
	ws := Wrappers{}
	ws = ws.Add(emptyWrapper)

	assert.Equal(t, len(ws), 1)
}

func TestWrappers_Wrap(t *testing.T) {
	ws := Wrappers{}

	ws = ws.Add(
		emptyWrapper,
		emptyWrapper,
	)
}

type counterHandler struct {
	calls int
}
