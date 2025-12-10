package main

import (
	"fmt"
	"strings"
)

func Test() {
	str := "  123"
	sp := strings.Split(str, "")
	fmt.Println(sp)
	fmt.Println(len(str))
}
