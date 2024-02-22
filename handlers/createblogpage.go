package handlers

import (
	"hsmyc/gocms/views/components/createblog"
	"hsmyc/gocms/views/layout"

	"github.com/a-h/templ"
)

func CreateBlogPageHandler() templ.Component {
	return layout.Index(nil, createblog.CreateBlogpage(), nil)
}
