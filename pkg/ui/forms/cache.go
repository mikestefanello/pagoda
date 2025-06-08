package forms

import (
	"net/http"

	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type Cache struct {
	CurrentValue string
	Value        string `form:"value"`
	form.Submission
}

func (f *Cache) Render(r *ui.Request) Node {
	return Form(
		ID("cache"),
		Method(http.MethodPost),
		Attr("hx-post", r.Path(routenames.CacheSubmit)),
		Card(CardParams{
			Title: "Test the cache",
			Body: Group{
				Span(Text("This route handler shows how the default in-memory cache works. Try updating the value using the form below and see how it persists after you reload the page.")),
				Span(Text("HTMX makes it easy to re-render the cached value after the form is submitted.")),
			},
			Color: ColorInfo,
			Size:  SizeMedium,
		}),
		Label(
			For("value"),
			Class("value"),
			Text("Value in cache: "),
		),
		If(f.CurrentValue != "", Badge(ColorSuccess, f.CurrentValue)),
		If(f.CurrentValue == "", Badge(ColorWarning, "empty")),
		InputField(InputFieldParams{
			Form:      f,
			FormField: "Value",
			Name:      "value",
			InputType: "text",
			Label:     "Value",
			Value:     f.Value,
		}),
		ControlGroup(
			FormButton(ColorPrimary, "Update cache"),
		),
		CSRF(r),
	)
}
