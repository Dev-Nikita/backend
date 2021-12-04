package main

import "fmt"

func getDataFromServer(resultCh chan string, serverName string) {
	resultCh <- "Data from server: " + serverName
}


func main() {
	
	res := make(chan string)

	go getDataFromServer(res, "Server1")
	time.Sleep(8 * time.Second)

	data := <- res
	fmt.Println(data)
}
