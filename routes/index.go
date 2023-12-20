package routes

import (
	"hsmyc/htmx/handlers"
	"net/http"

	"github.com/a-h/templ"
)

func Index() {
	http.Handle("/", templ.Handler(handlers.IndexHandler()))
}
