package main

import (
	"fmt"
	"hsmyc/gocms/configs"
	"hsmyc/gocms/routes"
	"net/http"
)

func main() {
	configs.ESBuilder()
	routes.InitializeContext()
	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
