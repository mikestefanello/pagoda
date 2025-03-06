package cache

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func TestCache_GetSet(t *testing.T) {
	key := "test"
	assert.Nil(t, Get(key))

	node := Div(Text("hello"))
	Set(key, node)

	got := Get(key)
	require.NotNil(t, got)

	// Check it was converted to a Raw component.
	_, ok := got.(NodeFunc)
	require.True(t, ok)

	// Both nodes should render the same string.
	buf1, buf2 := bytes.NewBuffer(nil), bytes.NewBuffer(nil)
	require.NoError(t, node.Render(buf1))
	require.NoError(t, got.Render(buf2))
	assert.Equal(t, buf1.String(), buf2.String())
}

func TestCache_SetIfNotExists(t *testing.T) {
	key := "test2"
	called := 0
	callback := func() Node {
		called++
		return Div(Text("hello"))
	}

	assertRender := func(n Node) {
		buf := bytes.NewBuffer(nil)
		require.NoError(t, n.Render(buf))
		assert.Equal(t, `<div>hello</div>`, buf.String())
	}

	got := SetIfNotExists(key, callback)
	assert.Equal(t, 1, called)
	require.NotNil(t, got)
	assertRender(got)

	got = SetIfNotExists(key, callback)
	assert.Equal(t, 1, called)
	require.NotNil(t, got)
	assertRender(got)
}
