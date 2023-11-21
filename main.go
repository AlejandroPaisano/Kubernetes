package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "my awsome go app")
}

func setupRoutes() {
	http.HandleFunc("/", homepage)
}

func main() {
	fmt.Println("go web app on  port 3001")
	setupRoutes()
	http.ListenAndServe(":3001", nil)
}
