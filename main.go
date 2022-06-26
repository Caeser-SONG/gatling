package main

import (
	"flag"
	"fmt"
)

func main() {
	var input string

	flag.StringVar(&input, "help", "help", "tools help")
	if input == "h" {
		fmt.Println("help ")
	}
	fmt.Println("ppp")
}
