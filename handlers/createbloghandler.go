package handlers

import (
	"context"

	"hsmyc/gocms/models"
	"time"
)

func CreateBlogHandler(blog models.Blogmodel) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	blogCollection.InsertOne(ctx, blog)
}
