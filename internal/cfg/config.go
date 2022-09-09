package cfg

import "os"

type Config struct {
	DataSource string `json:"data_source"`
}

var c *Config

func Get() *Config {
	if c == nil {
		var err error
		c, err = load()
		if err != nil {
			panic(err)
		}
	}

	return c
}

func load() (*Config, error) {
	_, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	c = &Config{
		DataSource: "ttt.db",
	}

	return c, nil
}
