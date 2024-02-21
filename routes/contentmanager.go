package routes

import (
	"hsmyc/htmx/handlers"
	"net/http"

	"github.com/a-h/templ"
)

func SingleType() {
	http.HandleFunc("/singletype/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(handlers.CreateSingleTypeHandler()).ServeHTTP(w, r)
	})

}
