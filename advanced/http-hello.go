package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Please visit http://127.0.0.1:12345/")
	http.HandleFunc("/", func(w http.ResponseWriter, req * http.Request) {
		var s = fmt.Sprintf("hello, world! -- Time: %s", time.Now().String())
		fmt.Fprintf(w, "%v\n", s)
		log.Printf("%v\n", s)
	})
	var err = http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}