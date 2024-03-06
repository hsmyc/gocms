package handlers

import (
	"hsmyc/gocms/views/layout"

	"github.com/a-h/templ"
)

func IndexHandler() templ.Component {

	return layout.Index(nil, nil, nil)
}
