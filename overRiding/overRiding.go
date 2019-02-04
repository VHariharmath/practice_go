package main

import "fmt"

type Father struct {
	first  string
	second string
}
type player struct {
	first      string
	second     string
	fatherName Father
	age        int
}

func (p player) playerinfo() string {
	return p.first + " " + p.second
}

func (p Father) playerinfo() string {
	return p.first + " " + p.second
}
func main() {
	p1 := player{
		first:  "Sachin",
		second: "Tendulkar",
		fatherName: Father{
			first:  "Ramesh",
			second: "Tendulkar",
		},
	}

	f1 := Father{
		first:  "abc",
		second: "xyz",
	}

	fmt.Println(p1.playerinfo())
	fmt.Println(f1.playerinfo())
	fmt.Println(p1.fatherName.playerinfo())

}
