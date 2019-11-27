package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	fmt.Println("Hello")

	nihon := &Person {
		Name:                 "Nihon",
		Age:                  30,
		SocialFollowers: &SocialFollowers{
			Youtube:              1500,
			Twitter:              2500,
		},
	}

	data, err := proto.Marshal(nihon)

	if err != nil {
		log.Fatal("Marshalling error: ", err)
	}

	fmt.Println(data)

	newNihon := &Person{}

	err = proto.Unmarshal(data, newNihon)

	if err != nil {
		log.Fatal("Unmarshal err ", err)
	}

	fmt.Println(newNihon.Age)
	fmt.Println(newNihon.Name)
	fmt.Println(newNihon.SocialFollowers)

}
