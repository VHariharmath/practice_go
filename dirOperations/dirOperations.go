package main

import (
	"fmt"
	"os"
)

func main() {

	dir, err := os.Open("../.")

	if err != nil {
		fmt.Println("failed to open dir")
		return
	}

	defer dir.Close()

	contents, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("failed to read dir")
		return
	}

	for _, v := range contents {
		fmt.Println(v)

	}
}
