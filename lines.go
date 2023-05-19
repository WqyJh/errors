package errors

type liner interface {
	line(stack bool) string
}

func Lines(err error, stack bool) []string {
	var errors = []string{}
	for err != nil {
		var line string
		switch err := err.(type) {
		case liner:
			line = err.line(stack)
		default:
			line = err.Error()
		}
		if len(line) != 0 {
			errors = append(errors, line)
		}
		err = Unwrap(err)
	}
	return errors
}
