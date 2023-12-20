package handlers

import (
	"hsmyc/htmx/db"
	"hsmyc/htmx/views/components/blogcard"
	"hsmyc/htmx/views/layout"

	"github.com/a-h/templ"
)

func IndexHandler() templ.Component {
	return layout.Index(nil, blogcard.Blogcard(db.GetBlogs()), nil)
}
