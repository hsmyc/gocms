package models

import (
	"github.com/a-h/templ"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: Create new structures for ReadingIndicator and Subheadings
type Blogmodel struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Title            string             `json:"title,omitempty" validate:"required"`
	Subtitle         string             `json:"subtitle,omitempty" validate:"required"`
	Author           string             `json:"author,omitempty" validate:"required"`
	Time             string             `json:"time,omitempty" validate:"required"`
	TldrHeader       string             `json:"tldrheader,omitempty" validate:"required"`
	Tldr             string             `json:"tldr,omitempty" validate:"required"`
	Content          string             `json:"content,omitempty" validate:"required"`
	Slug             templ.SafeURL      `json:"slug,omitempty" validate:"required"`
	Tags             []string           `json:"tags,omitempty" validate:"required"`
	Image            string             `json:"image,omitempty" validate:"required"`
	Links            []string           `json:"links,omitempty" validate:"required"`
	Subheadings      []string           `json:"subheadings,omitempty" validate:"required"`
	SubheadingLinks  []string           `json:"subheadinglinks,omitempty" validate:"required"`
	Topics           []string           `json:"topics,omitempty" validate:"required"`
	LikeCount        string             `json:"likecount,omitempty"`
	CommentCount     string             `json:"commentcount,omitempty"`
	ShareCount       string             `json:"sharecount,omitempty"`
	ReadingIndicator string             `json:"readingindicator,omitempty"`
}
