package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "jsdfaljfowjfa"
	sarry := strings.Split(s, ":")

	fmt.Println(sarry)
	fmt.Println(sarry[0])

	s = "jsdf:al:jfowjfa"
	sarry = strings.Split(s, ":")

	fmt.Println(sarry)
	fmt.Println(sarry[0])

}