package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/a-h/templ"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
func main() {
	comp := postTodo()
	http.Handle("/", templ.Handler(comp))
	http.HandleFunc("/osman", getHello)
	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
