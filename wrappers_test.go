package yac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Returns handler without any changes
var emptyWrapperFunc = WrapperFunc(func(h Handler) Handler {
	return h
})

func TestWrapper_Add(t *testing.T) {
	ws := Wrappers{}
	ws = ws.Add(emptyWrapperFunc)

	assert.Equal(t, len(ws), 1)
}

func TestWrappers_Wrap(t *testing.T) {
	ws := Wrappers{}

	thirdWrapper := &counterWrapper{}
	secondWrapper := &counterWrapper{nexts: []*counterWrapper{thirdWrapper}}
	firstWrapper := &counterWrapper{nexts: []*counterWrapper{secondWrapper, thirdWrapper}}

	ws = ws.Add(
		firstWrapper,
		secondWrapper,
		thirdWrapper,
	)

	ws.Wrap(emptyHandlerFunc)
	assert.Equal(t, uint32(1), firstWrapper.counter, "first has to be called first")
	assert.Equal(t, uint32(2), secondWrapper.counter, "second has to be called second")
	assert.Equal(t, uint32(3), thirdWrapper.counter, "third has to be called third")
}

// Wrapper that counts wraps
type counterWrapper struct {
	counter uint32
	nexts   []*counterWrapper
}

func (c *counterWrapper) Wrap(h Handler) Handler {
	for _, next := range c.nexts {
		next.counter++
	}

	c.counter++
	return h
}
