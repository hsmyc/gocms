package main

import (
	"fmt"
	"hsmyc/htmx/routes"
	"net/http"
)

func main() {
	routes.Blog()
	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
