package errors

type Options struct {
	FuncSep  string
	StackSep string
}

var (
	globalOptions Options = Options{
		FuncSep:  "\t",
		StackSep: "\n",
	}
)

func SetOptions(options Options) {
	globalOptions = options
}
