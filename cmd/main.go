package main

import (
	"fmt"
	"hsmyc/gocms/routes"
	"net/http"
)

func main() {
	routes.InitializeContext()
	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
