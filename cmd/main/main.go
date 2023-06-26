package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello")
	filerServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", filerServer)
	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
