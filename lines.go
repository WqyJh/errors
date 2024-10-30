package errors

type Liner interface {
	ErrorLine(stack bool) string
}

func Lines(err error, stack bool) []string {
	var errors = []string{}
	for err != nil {
		var line string
		switch err := err.(type) {
		case Liner:
			line = err.ErrorLine(stack)
		default:
			line = err.Error()
		}
		if len(line) > 0 {
			errors = append(errors, line)
		}
		err = Unwrap(err)
	}
	return errors
}
