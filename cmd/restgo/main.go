package main

import (
	"flag"
	"fmt"

	"github.com/damonchen/restgo"
)

var (
	file  string
	debug bool
)

func main() {
	flag.StringVar(&file, "file", "", "test file")
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.Parse()

	fmt.Printf("test file %s\n", file)
	err := restgo.TestFile(file, debug)
	if err != nil {
		fmt.Printf("test file error %s",
			err.Error())
	}
	return
}
