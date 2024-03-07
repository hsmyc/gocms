package routes

import (
	"hsmyc/gocms/handlers"
	"net/http"

	"github.com/a-h/templ"
)

func CreateSchema() {
	http.HandleFunc("/cms/schemas", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(handlers.SchemasHandler()).ServeHTTP(w, r)
	})
}
