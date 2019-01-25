package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	passwd := "welcome123"

	bs, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	loginpasswd := "welcome123"

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginpasswd))
	if err != nil {
		panic(err)
	}

	fmt.Println("YOU ARE LOGGED IN")

	loginpasswd = "welcome1234"

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginpasswd))
	if err != nil {
		fmt.Println("WRONG PASS WORD")
	}
}
