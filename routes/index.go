package routes

import (
	"hsmyc/htmx/views/layout"
	"net/http"

	"github.com/a-h/templ"
)

func Index() {
	index := layout.Index(nil, nil, nil)
	http.Handle("/", templ.Handler(index))
}
