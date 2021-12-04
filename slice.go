package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

)

func main()  {
	numberList := make([]int, 0, 3)

	var  scanInt string 

	for {
		fmt.Print("Enter a int number:")
		_, err := fmt.Scan(&scanInt)
		exit := strings.ToLower(scanInt)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		if exit == "x" {
			fmt.Println("Exit.")
			os.Exit(0)
		}

		parseIntScanInt, err := strconv.Atoi(scanInt)
		numberList = append(numberList, parseIntScanInt)
		sort.Ints(numberList)
		fmt.Println(numberList)

	}
}