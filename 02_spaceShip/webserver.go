package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", servHandle)
	fmt.Println("server online Port 8080")
	http.ListenAndServe(":8080", nil)
}

func servHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "wellecome to sapce")
}
