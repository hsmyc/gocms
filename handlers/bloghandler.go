package handlers

import (
	"hsmyc/htmx/views/components/blogpage"
	"hsmyc/htmx/views/layout"

	"github.com/a-h/templ"
)

func BlogHandler() templ.Component {
	t := blogpage.Title("Yuce's Blog")
	return layout.Index(t, nil, nil)
}
