package main

import "fmt"

type Career struct {
	avg   float64
	sr    float64
	cents int
}
type person struct {
	// Anonymous field, no field name
	Father
	first  string
	second string
	career Career
}

type Father struct {
	first  string
	second string
}

type playerName struct {
	first  string
	second string
}
type player struct {
	playerName
	age int
}

type emptyStruct struct {
}

func main() {
	p1 := person{
		Father: Father{
			first:  "xyz",
			second: "abc",
		},
		first:  "Virat",
		second: "Kohli",
		career: Career{
			avg:   59.6,
			sr:    95.5,
			cents: 39,
		},
	}
	fmt.Println(p1)

	fmt.Println(p1.first, p1.Father.first)

	p2 := player{
		playerName: playerName{
			first:  "Rohit",
			second: "Sharma",
		},
		age: 30,
	}

	// If "first" field name is not present in parent struct
	// then "first" named field from embedded struct can be used
	// directly like p2.first rather than p2.playerName.first
	// Below both prints same info
	fmt.Println(p2.first, p2.second, p2.age)
	fmt.Println(p2.playerName.first, p2.playerName.second, p2.age)

	// Anonymous struct, if struct is only used in only one place and not
	// require to declare will have this type of declaration and defination
	p3 := struct {
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

	fmt.Println(p3)
}
