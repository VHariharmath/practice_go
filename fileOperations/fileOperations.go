package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create file")
		return
	}

	defer file.Close()

	_, err = file.WriteString("Good morning Pune !")
	if err != nil {
		fmt.Println("Failed to write to file")
		return
	}

	// Set the file to start
	file.Seek(int64(os.SEEK_SET), io.SeekStart)

	bs, err := ioutil.ReadFile("test.txt")

	s := string(bs)

	fmt.Println(s)

}
