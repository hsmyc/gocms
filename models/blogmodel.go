package models

import "github.com/a-h/templ"

// ToDo: Create new struct for sub objects{Subheading,Author,Image}

type Blogmodel struct {
	Title           string
	Subtitle        string
	Author          string
	Time            string
	Tldr            string
	Content         string
	Url             templ.SafeURL
	Tags            []string
	Image           string
	Links           []string
	Subheadings     []string
	SubheadingLinks []string
}
