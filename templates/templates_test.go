package templates

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	_, err := Get().Open(fmt.Sprintf("pages/%s.gohtml", PageHome))
	require.NoError(t, err)
}

func TestGetOS(t *testing.T) {
	_, err := GetOS().Open(fmt.Sprintf("pages/%s.gohtml", PageHome))
	require.NoError(t, err)
}
