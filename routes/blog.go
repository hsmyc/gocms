package routers

import (
	"net/http"

	"github.com/a-h/templ"
)

func init() {

	t := title("Yuce's Blog")
	index := index(t, _, _)
	http.HandleFunc("/", templ.Handler(index))

}
