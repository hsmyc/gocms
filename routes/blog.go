package routes

import (
	"hsmyc/htmx/handlers"
	"net/http"

	"github.com/a-h/templ"
)

func Blog() {
	http.HandleFunc("/blog/", func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Path[6:]
		templ.Handler(handlers.BlogHandler(title)).ServeHTTP(w, r)
	})
}
