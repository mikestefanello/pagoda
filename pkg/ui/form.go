package ui

import (
	"github.com/mikestefanello/pagoda/pkg/form"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type (
	input struct {
		form      form.Form
		formField string
		name      string
		inputType string
		label     string
		value     string
	}

	radios struct {
		form      form.Form
		formField string
		name      string
		label     string
		value     string
		options   []radio
	}

	radio struct {
		value string
		label string
	}

	textarea struct {
		form      form.Form
		formField string
		name      string
		label     string
		value     string
	}
)

func formSubmit(label string) Node {
	return Div(
		Class("field is-grouped"),
		Div(
			Class("control"),
			Button(
				Class("button is-link"),
				Text(label),
			),
		),
	)
}

func formTextarea(el textarea) Node {
	return Div(
		Class("field"),
		Label(
			For("name"),
			Class("label"),
			Text(el.label),
		),
		Div(
			Class("control"),
			Textarea(
				ID(el.name),
				Name(el.name),
				Class("textarea "+formFieldStatusClass(el.form, el.formField)),
				Text(el.value),
			),
		),
		formFieldErrors(el.form, el.formField),
	)
}

func formRadios(el radios) Node {
	buttons := make(Group, 0, len(el.options))
	for _, opt := range el.options {
		buttons = append(buttons, Label(
			Class("radio"),
			Input(
				Type("radio"),
				Name(el.name),
				Value(opt.value),
				If(el.value == opt.value, Checked()),
			),
			Text(" "+opt.label),
		))
	}

	return Div(
		Class("control field"),
		Label(Class("label"), Text(el.label)),
		Div(
			Class("radios"),
			buttons,
		),
		formFieldErrors(el.form, el.formField),
	)
}

func formInput(el input) Node {
	return Div(
		Class("field"),
		Label(
			Class("label"),
			For(el.name),
			Text(el.label),
		),
		Div(
			Class("control"),
			Input(
				ID(el.name),
				Name(el.name),
				Type(el.inputType),
				Class("input "+formFieldStatusClass(el.form, el.formField)),
				Value(el.value),
			),
		),
		formFieldErrors(el.form, el.formField),
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

func csrf(r *request) Node {
	return Input(
		Type("hidden"),
		Name("csrf"),
		Value(r.CSRF),
	)
}
