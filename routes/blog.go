package routes

import (
	"hsmyc/htmx/handlers"
	"net/http"

	"github.com/a-h/templ"
)

func Blog() {
	index := handlers.BlogHandler()
	http.Handle("/blog", templ.Handler(index))
}
