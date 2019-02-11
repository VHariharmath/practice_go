package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	First      string
	Last       string
	Age        int
	unexported string
}

func main() {
	p1 := Person{"Virat", "Kohli", 29, "VK18"}
	json.NewEncoder(os.Stdout).Encode(p1)

	var p2 Person
	rdr := strings.NewReader(`{"First":"Ajinkya","Last":"Rahane","Age":30}`)
	json.NewDecoder(rdr).Decode(&p2)

	fmt.Println(p2.First)
	fmt.Println(p2.Last)
	fmt.Println(p2.Age)
}
