package handlers

import (
	"hsmyc/htmx/db"
	"hsmyc/htmx/views/components/blogpage"
	"hsmyc/htmx/views/layout"

	"github.com/a-h/templ"
)

func BlogHandler(blogTitle string) templ.Component {
	blogs := db.GetBlogs()
	for _, blog := range blogs {
		if blog.Title == blogTitle {
			return layout.Index(nil, blogpage.Blogpage(blog), nil)
		}
	}
	return layout.Index(nil, nil, nil)
}
