package handlers

import (
	"hsmyc/gocms/views/components/createschema"
	"hsmyc/gocms/views/layout"

	"github.com/a-h/templ"
)

func SchemasHandler() templ.Component {
	return layout.Index(nil, createschema.CreateSchema(), nil)
}
