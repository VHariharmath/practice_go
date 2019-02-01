package main

import "fmt"

func main() {
	m := map[string]int{
		"Virat":   29,
		"Ajinkya": 28,
		"Shami":   29,
	}

	s := []string{"Shikhar", "Rohit", "Virat", "Ajinkya", "MS", "Shami"}

	fmt.Println("Availability for the match")
	for _, item := range s {
		if _, ok := m[item]; ok {
			fmt.Printf("%v\t : Present in the squad and age = %v \n", item, m[item])
		} else {
			fmt.Printf("%v\t : Not Present in the squad\n", item)
		}
	}

	fmt.Println("Adding Manish and Karun to the team")
	m["Manish"] = 24
	m["Karun"] = 22

	fmt.Println("Squad now: ")
	for k, v := range m {
		fmt.Println(k, v)
	}
}
