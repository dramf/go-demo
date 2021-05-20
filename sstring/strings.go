package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Stroka 1"
	str2 := strings.Replace(str1, "1", "2", -1)
	fmt.Println("str1:", str1)
	fmt.Println("str2:", str2)
}