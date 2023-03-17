package main

import (
	"flag"
	"fmt"
)

func main() {

	var (
		help string
		url  string
		qps  int
	)
	flag.StringVar(&help, "help", "help", "tools help")
	flag.StringVar(&url, "url", "", "请求地址")
	flag.IntVar(&qps, "qps", 0, "qps值")
	flag.Parse()

	if help == "h" {
		fmt.Println("help ")
	}

	fmt.Println("ppp")

}
