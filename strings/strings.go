package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "My experience and knowledge need a “LANGUAGE” upgradation"
	fmt.Println(s)
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.Replace(s, " ", "-", -1))

}
