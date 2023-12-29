package handlers

import (
	"hsmyc/htmx/views/components/createblog"
	"hsmyc/htmx/views/layout"

	"github.com/a-h/templ"
)

func CreateBlogPageHandler() templ.Component {
	return layout.Index(nil, createblog.CreateBlogpage(), nil)
}
