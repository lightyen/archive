package main

import (
	"fmt"

	"github.com/lightyen/archive/zip"
)

func main() {
	err := zip.Compress("./test.zip", ".")
	if err != nil {
		fmt.Println(err)
	}
	err = zip.Extract("./data", "./test.zip")
	if err != nil {
		fmt.Println(err)
	}
}
