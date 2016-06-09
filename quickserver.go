package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var defaultPort = 12345

func getPort() int {

	if len(os.Args) > 1 {
		port := os.Args[1]
		portInt, _ := strconv.Atoi(port)
		return portInt
	}

	return defaultPort
}

func main() {
	port := getPort()
	wd, _ := os.Getwd()
	dir := http.Dir(wd)
	handler := http.FileServer(dir)
	fmt.Printf("quickserver starting on port %d\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), handler))
}
