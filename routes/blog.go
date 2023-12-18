package routes

import (
	"hsmyc/htmx/views/components/blogpage"
	"hsmyc/htmx/views/layout"
	"net/http"

	"github.com/a-h/templ"
)

func Blog() {
	t := blogpage.Title("Yuce's Blog")
	index := layout.Index(t, nil, nil)
	http.Handle("/blog", templ.Handler(index))
}
