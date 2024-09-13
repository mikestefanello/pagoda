package page

import (
	"fmt"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/tests"

	"github.com/stretchr/testify/assert"
)

func TestNewPager(t *testing.T) {
	e := echo.New()
	ctx, _ := tests.NewContext(e, "/")
	pgr := NewPager(ctx, 10)
	assert.Equal(t, 10, pgr.ItemsPerPage)
	assert.Equal(t, 1, pgr.Page)
	assert.Equal(t, 0, pgr.Items)
	assert.Equal(t, 1, pgr.Pages)

	ctx, _ = tests.NewContext(e, fmt.Sprintf("/abc?%s=%d", PageQueryKey, 2))
	pgr = NewPager(ctx, 10)
	assert.Equal(t, 2, pgr.Page)

	ctx, _ = tests.NewContext(e, fmt.Sprintf("/abc?%s=%d", PageQueryKey, -2))
	pgr = NewPager(ctx, 10)
	assert.Equal(t, 1, pgr.Page)
}

func TestPager_SetItems(t *testing.T) {
	ctx, _ := tests.NewContext(echo.New(), "/")
	pgr := NewPager(ctx, 20)
	pgr.SetItems(100)
	assert.Equal(t, 100, pgr.Items)
	assert.Equal(t, 5, pgr.Pages)

	pgr.SetItems(0)
	assert.Equal(t, 0, pgr.Items)
	assert.Equal(t, 1, pgr.Pages)
	assert.Equal(t, 1, pgr.Page)
}

func TestPager_IsBeginning(t *testing.T) {
	ctx, _ := tests.NewContext(echo.New(), "/")
	pgr := NewPager(ctx, 20)
	pgr.Pages = 10
	assert.True(t, pgr.IsBeginning())
	pgr.Page = 2
	assert.False(t, pgr.IsBeginning())
	pgr.Page = 1
	assert.True(t, pgr.IsBeginning())
}

func TestPager_IsEnd(t *testing.T) {
	ctx, _ := tests.NewContext(echo.New(), "/")
	pgr := NewPager(ctx, 20)
	pgr.Pages = 10
	assert.False(t, pgr.IsEnd())
	pgr.Page = 10
	assert.True(t, pgr.IsEnd())
	pgr.Page = 1
	assert.False(t, pgr.IsEnd())
}

func TestPager_GetOffset(t *testing.T) {
	ctx, _ := tests.NewContext(echo.New(), "/")
	pgr := NewPager(ctx, 20)
	assert.Equal(t, 0, pgr.GetOffset())
	pgr.Page = 2
	assert.Equal(t, 20, pgr.GetOffset())
	pgr.Page = 3
	assert.Equal(t, 40, pgr.GetOffset())
}
