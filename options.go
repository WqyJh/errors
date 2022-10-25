package errors

type Config struct {
	FuncSep  string
	StackSep string
	MsgSep   string
}

type Option func(*Config)

var (
	globalOptions = Config{
		FuncSep:  "\t",
		StackSep: "\n",
		MsgSep:   " : ",
	}
)

func WithFuncSep(sep string) Option {
	return func(c *Config) {
		c.FuncSep = sep
	}
}

func WithStackSep(sep string) Option {
	return func(c *Config) {
		c.StackSep = sep
	}
}

func WithMsgSep(sep string) Option {
	return func(c *Config) {
		c.MsgSep = sep
	}
}

func SetOptions(options ...Option) {
	for _, option := range options {
		option(&globalOptions)
	}
}
