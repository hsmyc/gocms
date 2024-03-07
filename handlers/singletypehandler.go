package handlers

import (
	"fmt"
	"hsmyc/gocms/models"
	"hsmyc/gocms/views/components/createschema/singletype"
	"hsmyc/gocms/views/layout"

	"github.com/a-h/templ"
)

func SingleTypeCreatePageHandler() templ.Component {
	return layout.Index(nil, singletype.CreateSingleType(), nil)
}

func CreateSingleTypeHandler(Doc models.SingleTypeModel) {
	data := models.SingleTypeModel{
		Name:   Doc.Name,
		Fields: Doc.Fields,
	}
	//TODO: Save data to JSON file

	fmt.Println(data)
}
