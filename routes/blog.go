package routes

import (
	"hsmyc/htmx/components"
	"hsmyc/htmx/handlers"
	"net/http"

	"github.com/a-h/templ"
)

func Blog() {
	t := components.Title("Yuce's Blog")
	index := handlers.Index(t, nil, nil)
	http.Handle("/", templ.Handler(index))
}
