package templates

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	_, err := Get().Open("pages/home.gohtml")
	require.NoError(t, err)
}

func TestGetOS(t *testing.T) {
	_, err := GetOS().Open("pages/home.gohtml")
	require.NoError(t, err)
}
