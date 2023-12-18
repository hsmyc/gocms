package routes

import (
	"hsmyc/htmx/views/components/navbar"
	"hsmyc/htmx/views/layout"
	"net/http"

	"github.com/a-h/templ"
)

func Index() {
	var navbarMenu []string
	navbarMenu = append(navbarMenu, "Home")
	navbarMenu = append(navbarMenu, "Blog")
	navbarMenu = append(navbarMenu, "About")
	nav := navbar.Navbar(navbarMenu)
	index := layout.Index(nav, nil, nil)
	http.Handle("/", templ.Handler(index))
}
