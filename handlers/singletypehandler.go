package handlers

import (
	"hsmyc/htmx/views/components/contentschema/singletype"
	"hsmyc/htmx/views/layout"

	"github.com/a-h/templ"
)

func CreateSingleTypeHandler() templ.Component {
	return layout.Index(nil, singletype.CreateSingleType(), nil)
}
