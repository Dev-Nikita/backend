package main

import (
	"fmt"
	"strings"

)

func main() {
	var str string
	
	fmt.Println("Enter a word ")
	fmt.Scan(&str)
	str = strings.ToLower(str)
	fmt.Printf("word is %s.\n", str)

	lenStr := len(str)
	if  lenStr < 1 {
		fmt.Println("NOT FOUND")
		return
	}

	if strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str, "a") {
		fmt.Println("FOUND!")
	} else {
		fmt.Println("NOT FOUND!")
	}

}