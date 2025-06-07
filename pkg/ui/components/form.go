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

	OptionsParams struct {
		Form      form.Form
		FormField string
		Name      string
		Label     string
		Value     string
		Options   []Choice
	}

	Choice struct {
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

	CheckboxParams struct {
		Form      form.Form
		FormField string
		Name      string
		Label     string
		Checked   bool
	}
)

func ControlGroup(controls ...Node) Node {
	return Div(
		Class("mt-2 flex gap-2"),
		Group(controls),
	)
}

func TextareaField(el TextareaFieldParams) Node {
	return FieldSet(
		Class("fieldset"),
		If(len(el.Label) > 0, Legend(
			Class("fieldset-legend"),
			Text(el.Label),
		)),
		Textarea(
			Class("textarea h-24 "+formFieldStatusClass(el.Form, el.FormField)),
			ID(el.Name),
			Name(el.Name),
			Text(el.Value),
		),
		Help(el.Help),
		formFieldErrors(el.Form, el.FormField),
	)
}

func Radios(el OptionsParams) Node {
	buttons := make(Group, len(el.Options))
	for i, opt := range el.Options {
		id := "radio-" + el.Name + "-" + opt.Value
		buttons[i] = Div(
			Class("mb-2"),
			Input(
				ID(id),
				Type("radio"),
				Name(el.Name),
				Value(opt.Value),
				Class("radio mr-1 "+formFieldStatusClass(el.Form, el.FormField)),
				If(el.Value == opt.Value, Checked()),
			),
			Label(
				Text(opt.Label),
				For(id),
			),
		)
	}

	return FieldSet(
		If(len(el.Label) > 0, Legend(
			Class("fieldset-legend"),
			Text(el.Label),
		)),
		buttons,
		formFieldErrors(el.Form, el.FormField),
	)
}

func SelectList(el OptionsParams) Node {
	buttons := make(Group, len(el.Options))
	for i, opt := range el.Options {
		buttons[i] = Option(
			Text(opt.Label),
			Value(opt.Value),
			If(opt.Value == el.Value, Attr("selected")),
		)
	}

	return Div(
		Class("control field"),
		Label(Class("label"), Text(el.Label)),
		Div(
			Class("select"),
			Select(
				Name(el.Name),
				buttons,
			),
		),
		formFieldErrors(el.Form, el.FormField),
	)
}

func Checkbox(el CheckboxParams) Node {
	return Div(
		Label(
			Class("label"),
			Input(
				Class("checkbox"),
				Type("checkbox"),
				Name(el.Name),
				If(el.Checked, Checked()),
				Value("true"),
			),
			Text(" "+el.Label),
		),
		formFieldErrors(el.Form, el.FormField),
	)
}

func InputField(el InputFieldParams) Node {
	return FieldSet(
		Class("fieldset"),
		If(len(el.Label) > 0, Legend(
			Class("fieldset-legend"),
			Text(el.Label),
		)),
		Input(
			ID(el.Name),
			Name(el.Name),
			Type(el.InputType),
			Class("input "+formFieldStatusClass(el.Form, el.FormField)),
			Value(el.Value),
			If(el.Placeholder != "", Placeholder(el.Placeholder)),
		),
		Help(el.Help),
		formFieldErrors(el.Form, el.FormField),
	)
}

func Help(text string) Node {
	return If(len(text) > 0, Div(
		Class("label"),
		Text(text),
	))
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
	case fm == nil:
		return ""
	case !fm.IsSubmitted():
		return ""
	case fm.FieldHasErrors(formField):
		return "input-error"
	default:
		return "input-success"
	}
}

func formFieldErrors(fm form.Form, field string) Node {
	if fm == nil {
		return nil
	}

	errs := fm.GetFieldErrors(field)
	if len(errs) == 0 {
		return nil
	}

	g := make(Group, len(errs))
	for i, err := range errs {
		g[i] = Div(
			Class("text-error"),
			Text(err),
		)
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

func FormButton(color Color, label string) Node {
	return Button(
		Class("btn "+buttonColor(color)),
		Text(label),
	)
}

func ButtonLink(color Color, href, label string) Node {
	return A(
		Href(href),
		Class("btn "+buttonColor(color)),
		Text(label),
	)
}

func buttonColor(color Color) string {
	// Only colors being used are included so unused styles are not compiled.
	switch color {
	case ColorPrimary:
		return "btn-primary"
	case ColorInfo:
		return "btn-info"
	case ColorAccent:
		return "btn-accent"
	case ColorError:
		return "btn-error"
	case ColorLink:
		return "btn-link"
	default:
		return ""
	}
}
