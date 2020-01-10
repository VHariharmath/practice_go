package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isValidIP(ipAddr string) bool {

	strSlice := strings.Split(ipAddr, ".")

	if len(strSlice) < 4 {
		fmt.Println("only", len(strSlice), "present")
		return false
	}

	for _, x := range strSlice {
		a, err := strconv.Atoi(x)
		if err != nil {
			fmt.Println(x, "is not a number")
			return false
		}

		if a >= 0 && a <= 255 {
			continue
		}

	}

	return true
}

func main() {

	addresses := []string{"192.168.0.1", "192.168.0", "goh", "", "10.0.5.25", "125.512.100.abc"}

	for _, ipAddr := range addresses {

		if isValidIP(ipAddr) {
			fmt.Println("Valid IP")
		} else {
			fmt.Println("Invalid IP")
		}

	}

}
