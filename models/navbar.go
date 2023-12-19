package models

import "github.com/a-h/templ"

type NavbarItem struct {
	Title string
	Url   templ.SafeURL
}
