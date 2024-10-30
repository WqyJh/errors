package errors

import "fmt"

type ApiConfig struct {
	CallerSkip int
}

type errorsApi struct {
	cfg ApiConfig
}

func NewErrorsApi(cfg ApiConfig) *errorsApi {
	return &errorsApi{
		cfg: cfg,
	}
}

var globalErrorsApi = NewErrorsApi(ApiConfig{
	CallerSkip: 2,
})

func (e *errorsApi) New(message string) error {
	return &fundamental{
		msg:   message,
		stack: callers(e.cfg.CallerSkip),
	}
}

func (e *errorsApi) Errorf(format string, args ...interface{}) error {
	return &fundamental{
		msg:   fmt.Sprintf(format, args...),
		stack: callers(e.cfg.CallerSkip),
	}
}

func (e *errorsApi) WithStack(err error) error {
	if err == nil {
		return nil
	}
	return &withStack{
		withMessage{
			cause: err,
			msg:   "",
		},
		callers(e.cfg.CallerSkip),
	}
}

func (e *errorsApi) Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return &withStack{
		withMessage{
			cause: err,
			msg:   message,
		},
		callers(e.cfg.CallerSkip),
	}
}

func (e *errorsApi) Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &withStack{
		withMessage{
			cause: err,
			msg:   fmt.Sprintf(format, args...),
		},
		callers(e.cfg.CallerSkip),
	}
}

func (e *errorsApi) WithMessage(err error, message string) error {
	if err == nil {
		return nil
	}
	return &withMessage{
		cause: err,
		msg:   message,
	}
}

func (e *errorsApi) WithMessagef(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &withMessage{
		cause: err,
		msg:   fmt.Sprintf(format, args...),
	}
}

func (e *errorsApi) WithDetails(err error, details ...any) error {
	if err == nil {
		return nil
	}
	return &withDetails{
		cause:   err,
		details: details,
	}
}

func (e *errorsApi) Details(err error) ([]any, bool) {
	if w, ok := err.(*withDetails); ok {
		return w.details, true
	}
	return nil, false
}
