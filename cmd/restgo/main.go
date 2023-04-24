package main

import (
	"flag"
	"fmt"
	"github.com/damonchen/restgo"
)

var (
	file string
)

func main() {
	flag.StringVar(&file, "file", "", "test file")
	flag.Parse()

	fmt.Printf("test file %s\n", file)
	err := restgo.TestFile(file)
	if err != nil {
		fmt.Printf("test file error %s",
			err.Error())
	}
	return
}
