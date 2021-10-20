package variableparameters

type Config struct {
	id      int
	name    string
	address string
}

type Options func(config *Config)

func WithIdOption(id int) Options {
	return func(config *Config) {
		config.id = id
	}
}

func WithNameOption(name string) Options {
	return func(config *Config) {
		config.name = name
	}
}

func WithaddressOption(address string) Options {
	return func(config *Config) {
		config.address = address
	}
}

func NewOptions(options ...Options) *Config {
	config := &Config{}

	for _, option := range options {
		option(config)
	}

	return config
}
