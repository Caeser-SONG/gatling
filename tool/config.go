package tool

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

var (
	one  sync.Once
	conf Conf
)

type Conf struct {
	Suit Suit
}

type Suit struct {
	Groups []Group `toml:"group"`
}

type Group struct {
	Method string `toml:"method"`
	Url    string `toml:"url"`
	Data   string `toml:"data"`
	Header string `toml:"header"`
}

func Config() *Conf {
	//conf = new(Suit)
	one.Do(func() {

		filepath, err := filepath.Abs("../conf/press.toml")
		fmt.Println(filepath)
		if err != nil {
			panic(err)

		}
		if _, err := toml.DecodeFile(filepath, &conf); err != nil {
			panic(err)
		}
	})
	return &conf
}
