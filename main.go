package main

import (
	"fmt"

	"github.com/lightyen/archive/zip"
)

func main() {
	if err := zip.Compress("./test.zip", "."); err != nil {
		fmt.Println(err)
	}
	if err := zip.Extract("./test", "./test.zip"); err != nil {
		fmt.Println(err)
	}
}
