package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First      string
	Second     string
	Age        int `json:"Players age"`
	unexported int
}

type player struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{"Virat", "Kohli", 29, 59}
	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(bs)
	fmt.Println(string(bs))

	// Without tags
	bs = []byte(`{"First":"Virat", "Last":"Kohli", "Age":30}`)

	var p2 player
	json.Unmarshal(bs, &p2)

	fmt.Println(p2.First)
	fmt.Println(p2.Last)
	fmt.Println(p2.Age)

	// with tags
	bs = []byte(`{"First":"Virat", "Second":"Kohli", "Players age":31}`)
	var p3 person
	json.Unmarshal(bs, &p3)
	fmt.Println(p3.First)
	fmt.Println(p3.Second)
	fmt.Println(p3.Age)
}
