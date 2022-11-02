package controller

import (
	"testing"

	"github.com/mikestefanello/pagoda/pkg/tests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormSubmission(t *testing.T) {
	type formTest struct {
		Name       string `validate:"required"`
		Email      string `validate:"required,email"`
		Submission FormSubmission
	}

	ctx, _ := tests.NewContext(c.Web, "/")
	form := formTest{
		Name:  "",
		Email: "a@a.com",
	}
	err := form.Submission.Process(ctx, form)
	assert.NoError(t, err)

	assert.True(t, form.Submission.HasErrors())
	assert.True(t, form.Submission.FieldHasErrors("Name"))
	assert.False(t, form.Submission.FieldHasErrors("Email"))
	require.Len(t, form.Submission.GetFieldErrors("Name"), 1)
	assert.Len(t, form.Submission.GetFieldErrors("Email"), 0)
	assert.Equal(t, "This field is required.", form.Submission.GetFieldErrors("Name")[0])
	assert.Equal(t, "is-danger", form.Submission.GetFieldStatusClass("Name"))
	assert.Equal(t, "is-success", form.Submission.GetFieldStatusClass("Email"))
	assert.False(t, form.Submission.IsDone())
}
