package msg

import (
	"testing"

	"goweb/tests"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"

	"k8s.io/apimachinery/pkg/util/rand"

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

	text := rand.String(10)
	Success(ctx, text)
	assertMsg(TypeSuccess, text)

	text = rand.String(10)
	Info(ctx, text)
	assertMsg(TypeInfo, text)

	text = rand.String(10)
	Danger(ctx, text)
	assertMsg(TypeDanger, text)

	text = rand.String(10)
	Warning(ctx, text)
	assertMsg(TypeWarning, text)

	text = rand.String(10)
	Set(ctx, TypeSuccess, text)
	assertMsg(TypeSuccess, text)
}
