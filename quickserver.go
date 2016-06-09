package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := 12345
	wd, _ := os.Getwd()
	dir := http.Dir(wd)
	handler := http.FileServer(dir)
	fmt.Printf("quickserver starting on port %d\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), handler))
}
