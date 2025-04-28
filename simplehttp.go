package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Привет Мир версия 4 !")
	})
	log.Fatal(http.ListenAndServe(":8082", nil))
}
