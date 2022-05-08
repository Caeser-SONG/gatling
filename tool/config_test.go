package tool

import (
	"fmt"
	"testing"
)

func TestConf(t *testing.T) {
	a := Config()
	fmt.Println(a)
}
