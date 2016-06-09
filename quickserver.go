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

func getDir() string {
	wd, _ := os.Getwd()
	return wd
}

func main() {
	port := getPort()
	dir := http.Dir(getDir())
	handler := http.FileServer(dir)

	fmt.Printf("quickserver starting at http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), handler))
}
