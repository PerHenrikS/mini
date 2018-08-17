package dev

import (
	"fmt"
	"net/http"
)

//Serve : works as a temporary server for dev purposes
func Serve() {
	http.Handle("/", http.FileServer(http.Dir("./webpage")))

	fmt.Println("Listening at port 8000")

	http.ListenAndServe(":8000", nil)
}
