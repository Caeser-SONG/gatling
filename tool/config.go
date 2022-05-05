package tool

import (
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

var (
	one  sync.Once
	conf *Suit
)

type Suit struct {
	groups []Group
}

type Group struct {
	Url    string `toml:"url"`
	Data   string `toml:"data"`
	Header string `toml:"header"`
}

func Config() *Suit {
	one.Do(func() {
		filepath, err := filepath.Abs("./press.toml")
		if err != nil {
			panic(err)

		}
		if _, err := toml.DecodeFile(filepath, conf); err != nil {
			panic(err)
		}
	})
	return conf
}
