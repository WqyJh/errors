package errors

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLines(t *testing.T) {
	tests := []struct {
		lines []string
		err   error
	}{
		{[]string{}, nil},
		{[]string{"foo"}, fmt.Errorf("foo")},
		{[]string{"foo\ngithub.com/pkg/errors.TestLines\t.+/github.com/pkg/errors/lines_test.go:18"}, New("foo")},
		{[]string{
			"bar\ngithub.com/pkg/errors.TestLines\t.+/github.com/pkg/errors/lines_test.go:22",
			"foo\ngithub.com/pkg/errors.TestLines\t.+/github.com/pkg/errors/lines_test.go:22",
		}, Wrap(New("foo"), "bar")},
		{[]string{
			"github.com/pkg/errors.TestLines\t.+/github.com/pkg/errors/lines_test.go:26",
			"foo\ngithub.com/pkg/errors.TestLines\t.+/github.com/pkg/errors/lines_test.go:26",
		}, WithStack(New("foo"))},
		{[]string{
			"bar 1\ngithub.com/pkg/errors.TestLines\t.+/github.com/pkg/errors/lines_test.go:31",
			"github.com/pkg/errors.TestLines\t.+/github.com/pkg/errors/lines_test.go:31",
			"foo\ngithub.com/pkg/errors.TestLines\t.+/github.com/pkg/errors/lines_test.go:31",
		}, Wrapf(WithStack(New("foo")), "bar %d", 1)},
		{[]string{
			"bar",
			"github.com/pkg/errors.TestLines\t.+/github.com/pkg/errors/lines_test.go:36",
			"foo\ngithub.com/pkg/errors.TestLines\t.+/github.com/pkg/errors/lines_test.go:36",
		}, WithMessage(WithStack(New("foo")), "bar")},
		{[]string{
			"bar 1",
			"github.com/pkg/errors.TestLines\t.+/github.com/pkg/errors/lines_test.go:41",
			"foo\ngithub.com/pkg/errors.TestLines\t.+/github.com/pkg/errors/lines_test.go:41",
		}, WithMessagef(WithStack(New("foo")), "bar %d", 1)},
		{[]string{
			"bar 1",
			"github.com/pkg/errors.TestLines\t.+/github.com/pkg/errors/lines_test.go:46",
			"foo 1\ngithub.com/pkg/errors.TestLines\t.+/github.com/pkg/errors/lines_test.go:46",
		}, Wrapf(WithStack(Errorf("foo %d", 1)), "bar %d", 1)},
	}

	for _, tt := range tests {
		got := Lines(tt.err, true)
		matchLines(t, tt.lines, got)
	}
}

func matchLines(t *testing.T, want []string, got []string) {
	assert.Equal(t, len(want), len(got))
	for i, w := range want {
		match, err := regexp.MatchString(w, got[i])
		assert.NoError(t, err, "regexp.MatchString(%q, %q)", w, got[i])
		assert.True(t, match, "regexp.MatchString(%q, %q)", w, got[i])
	}
}

type codeError struct {
	Code int
	Msg  string
}

func NewCodeError(code int, msg string) *codeError {
	return &codeError{Code: code, Msg: msg}
}

func (e *codeError) Error() string {
	return fmt.Sprintf("code error: %d, %s", e.Code, e.Msg)
}

func TestLinesNoStack(t *testing.T) {
	tests := []struct {
		lines []string
		err   error
	}{
		{[]string{}, nil},
		{[]string{"foo"}, fmt.Errorf("foo")},
		{[]string{"foo"}, New("foo")},
		{[]string{
			"bar",
			"foo",
		}, Wrap(New("foo"), "bar")},
		{[]string{
			"foo",
		}, WithStack(New("foo"))},
		{[]string{
			"bar 1",
			"foo",
		}, Wrapf(WithStack(New("foo")), "bar %d", 1)},
		{[]string{
			"bar",
			"foo",
		}, WithMessage(WithStack(New("foo")), "bar")},
		{[]string{
			"bar 1",
			"foo",
		}, WithMessagef(WithStack(New("foo")), "bar %d", 1)},
		{[]string{
			"bar 1",
			"foo 1",
		}, Wrapf(WithStack(Errorf("foo %d", 1)), "bar %d", 1)},
		{[]string{
			"code error: 1, foo",
		}, NewCodeError(1, "foo")},
		{[]string{
			"bar",
			"code error: 1, foo",
		}, Wrap(NewCodeError(1, "foo"), "bar")},
	}

	for _, tt := range tests {
		got := Lines(tt.err, false)
		assert.Equal(t, tt.lines, got)
	}
}
