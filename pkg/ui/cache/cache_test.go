package cache

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func TestCache(t *testing.T) {
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
