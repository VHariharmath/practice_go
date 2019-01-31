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
			fmt.Printf("%v\t : Available \n", item)
		} else {
			fmt.Printf("%v\t : Not available \n", item)
		}
	}
}