package routes

import (
	"hsmyc/htmx/handlers"
	"hsmyc/htmx/models"
	"net/http"
	"strings"

	"github.com/a-h/templ"
)

var blog models.Blogmodel

func Blog() {
	http.HandleFunc("/blog/", func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Path[6:]
		templ.Handler(handlers.BlogHandler(title)).ServeHTTP(w, r)
	})
}

func CreateBlog() {
	http.HandleFunc("/createblog", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			templ.Handler(handlers.CreateBlogPageHandler()).ServeHTTP(w, r)
		} else if r.Method == http.MethodPost {

			err := r.ParseForm()

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			blog.Title = r.FormValue("title")
			blog.Subtitle = r.FormValue("subtitle")
			blog.Author = r.FormValue("author")
			blog.Time = r.FormValue("time")
			blog.TldrHeader = r.FormValue("tldrheader")
			blog.Tldr = r.FormValue("tldr")
			blog.Content = r.FormValue("content")
			blog.Slug = templ.SafeURL(r.FormValue("slug"))
			blog.Image = r.FormValue("image")

			blog.Tags = strings.Split(r.FormValue("tags"), ",")
			blog.Links = strings.Split(r.FormValue("links"), ",")
			blog.Subheadings = strings.Split(r.FormValue("subheadings"), ",")
			blog.SubheadingLinks = strings.Split(r.FormValue("subheadinglinks"), ",")
			blog.Topics = strings.Split(r.FormValue("topics"), ",")

			handlers.CreateBlogHandler(blog)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Blog created successfully"))
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

}
