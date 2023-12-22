package handlers

import (
	"context"
	"hsmyc/htmx/models"
	"hsmyc/htmx/views/components/blogcard"
	"hsmyc/htmx/views/layout"
	"time"

	"github.com/a-h/templ"
	"go.mongodb.org/mongo-driver/bson"
)

func IndexHandler() templ.Component {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var blogs []models.Blogmodel
	cur, err := blogCollection.Find(ctx, bson.M{})
	println(cur)
	if err != nil {
		return layout.Index(nil, nil, nil)
	}
	err = cur.Decode(&blogs)

	if err != nil {
		return layout.Index(nil, nil, nil)
	}
	return layout.Index(nil, blogcard.Blogcard(blogs), nil)
}
