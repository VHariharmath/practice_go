package main

import "fmt"

type Career struct {
	avg   float64
	sr    float64
	cents int
}
type person struct {
	first  string
	second string
	career Career
}

type emptyStruct struct {
}

func main() {
	p1 := person{
		first:  "Virat",
		second: "Kohli",
		career: Career{
			avg:   59.6,
			sr:    95.5,
			cents: 39,
		},
	}
	fmt.Println(p1)

	// Anonymous struct
	p2 := struct {
		first  string
		second string
		career Career
	}{
		first:  "Ajinkya",
		second: "Rahane",
		career: struct {
			avg   float64
			sr    float64
			cents int
		}{
			avg:   45,
			sr:    85,
			cents: 15,
		},
	}

	fmt.Println(p2)
}
