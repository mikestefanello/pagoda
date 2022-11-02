package msg

import (
	"testing"

	"github.com/mikestefanello/pagoda/pkg/tests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/labstack/echo/v4"
)

func TestMsg(t *testing.T) {
	e := echo.New()
	ctx, _ := tests.NewContext(e, "/")
	tests.InitSession(ctx)

	assertMsg := func(typ Type, message string) {
		ret := Get(ctx, typ)
		require.Len(t, ret, 1)
		assert.Equal(t, message, ret[0])
		ret = Get(ctx, typ)
		require.Len(t, ret, 0)
	}

	text := "aaa"
	Success(ctx, text)
	assertMsg(TypeSuccess, text)

	text = "bbb"
	Info(ctx, text)
	assertMsg(TypeInfo, text)

	text = "ccc"
	Danger(ctx, text)
	assertMsg(TypeDanger, text)

	text = "ddd"
	Warning(ctx, text)
	assertMsg(TypeWarning, text)

	text = "eee"
	Set(ctx, TypeSuccess, text)
	assertMsg(TypeSuccess, text)
}
