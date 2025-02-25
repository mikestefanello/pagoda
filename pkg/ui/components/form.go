package components

import (
	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type (
	InputFieldParams struct {
		Form        form.Form
		FormField   string
		Name        string
		InputType   string
		Label       string
		Value       string
		Placeholder string
		Help        string
	}

	RadiosParams struct {
		Form      form.Form
		FormField string
		Name      string
		Label     string
		Value     string
		Options   []Radio
	}

	Radio struct {
		Value string
		Label string
	}

	TextareaFieldParams struct {
		Form      form.Form
		FormField string
		Name      string
		Label     string
		Value     string
		Help      string
	}
)

func ControlGroup(controls ...Node) Node {
	g := make(Group, 0, len(controls))
	for _, control := range controls {
		g = append(g, Div(
			Class("control"),
			control,
		))
	}

	return Div(
		Class("field is-grouped"),
		g,
	)
}

func TextareaField(el TextareaFieldParams) Node {
	return Div(
		Class("field"),
		Label(
			For("name"),
			Class("label"),
			Text(el.Label),
		),
		Div(
			Class("control"),
			Textarea(
				ID(el.Name),
				Name(el.Name),
				Class("textarea "+formFieldStatusClass(el.Form, el.FormField)),
				Text(el.Value),
			),
		),
		If(el.Help != "", P(Class("help"), Text(el.Help))),
		formFieldErrors(el.Form, el.FormField),
	)
}

func Radios(el RadiosParams) Node {
	buttons := make(Group, 0, len(el.Options))
	for _, opt := range el.Options {
		buttons = append(buttons, Label(
			Class("radio"),
			Input(
				Type("radio"),
				Name(el.Name),
				Value(opt.Value),
				If(el.Value == opt.Value, Checked()),
			),
			Text(" "+opt.Label),
		))
	}

	return Div(
		Class("control field"),
		Label(Class("label"), Text(el.Label)),
		Div(
			Class("radios"),
			buttons,
		),
		formFieldErrors(el.Form, el.FormField),
	)
}

func InputField(el InputFieldParams) Node {
	return Div(
		Class("field"),
		Label(
			Class("label"),
			For(el.Name),
			Text(el.Label),
		),
		Div(
			Class("control"),
			Input(
				ID(el.Name),
				Name(el.Name),
				Type(el.InputType),
				If(el.Placeholder != "", Placeholder(el.Placeholder)),
				Class("input "+formFieldStatusClass(el.Form, el.FormField)),
				Value(el.Value),
			),
		),
		If(el.Help != "", P(Class("help"), Text(el.Help))),
		formFieldErrors(el.Form, el.FormField),
	)
}

func FileField(name, label string) Node {
	return Div(
		Class("field file"),
		Label(
			Class("file-label"),
			Input(
				Class("file-input"),
				Type("file"),
				Name(name),
			),
			Span(
				Class("file-cta"),
				Span(
					Class("file-label"),
					Text(label),
				),
			),
		),
	)
}

func formFieldStatusClass(fm form.Form, formField string) string {
	switch {
	case !fm.IsSubmitted():
		return ""
	case fm.FieldHasErrors(formField):
		return "is-danger"
	default:
		return "is-success"
	}
}

func formFieldErrors(fm form.Form, field string) Node {
	errs := fm.GetFieldErrors(field)
	if len(errs) == 0 {
		return nil
	}

	g := make(Group, 0, len(errs))
	for _, err := range errs {
		g = append(g, P(
			Class("help is-danger"),
			Text(err),
		))
	}

	return g
}

func CSRF(r *ui.Request) Node {
	return Input(
		Type("hidden"),
		Name("csrf"),
		Value(r.CSRF),
	)
}

func FormButton(class, label string) Node {
	return Button(
		Class("button "+class),
		Text(label),
	)
}

func ButtonLink(href, class, label string) Node {
	return A(
		Href(href),
		Class("button "+class),
		Text(label),
	)
}
