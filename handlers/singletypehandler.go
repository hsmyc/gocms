package handlers

import (
	"context"
	"fmt"
	"hsmyc/gocms/db"
	"hsmyc/gocms/models"
	"hsmyc/gocms/views/components/createschema/singletype"
	"hsmyc/gocms/views/layout"
	"time"

	"github.com/a-h/templ"
)

var DB = db.ConnectDB()
var SingleTypeCollection = db.GetCollection(DB, "singleType")

func SingleTypeCreatePageHandler() templ.Component {
	return layout.Index(nil, singletype.CreateSingleType(), nil)
}

func CreateSingleTypeHandler(Doc models.SingleTypeModel) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	data := models.SingleTypeModel{
		Name:   Doc.Name,
		Fields: Doc.Fields,
	}
	_, err := SingleTypeCollection.InsertOne(ctx, data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("SingleType Created")

}
