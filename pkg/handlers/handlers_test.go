package handlers

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSetHandlers(t *testing.T) {
	handlers = []Handler{}
	assert.Empty(t, GetHandlers())
	h := new(Pages)
	Register(h)
	got := GetHandlers()
	require.Len(t, got, 1)
	assert.Equal(t, h, got[0])
}

func TestFail(t *testing.T) {
	err := fail(errors.New("err message"), "log message")
	require.IsType(t, new(echo.HTTPError), err)
	he := err.(*echo.HTTPError)
	assert.Equal(t, http.StatusInternalServerError, he.Code)
	assert.Equal(t, "log message: err message", he.Message)
}
