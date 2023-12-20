package db

import "hsmyc/htmx/models"

func GetBlogs() []models.Blogmodel {
	var blogs []models.Blogmodel
	blogs = append(blogs, models.Blogmodel{
		Title:           "Title1",
		Subtitle:        "Subtitle",
		Author:          "Author",
		Time:            "02-01-1991",
		Tldr:            "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla vitae elit libero, a pharetra augue. Donec id elit non mi porta gravida at eget metus. Nullam id dolor id nibh ultricies vehicula ut id elit.",
		Content:         "Lorem ipsum dolor sit AhMET, consectetur adipiscing elit. Nulla vitae elit libero, a pharetra augue. Donec id elit non mi porta gravida at eget metus. Nullam id dolor id nibh ultricies vehicula ut id elit.",
		Url:             "blog/Title1",
		Tags:            []string{"tag1", "tag2"},
		Image:           "https://picsum.photos/200/300",
		Links:           []string{"link1", "link2"},
		Subheadings:     []string{"sub1", "sub2"},
		SubheadingLinks: []string{"sublink1", "sublink2"},
	})
	blogs = append(blogs, models.Blogmodel{
		Title:           "Title2",
		Subtitle:        "Subtitle2",
		Author:          "Author2",
		Time:            "02-01-1991",
		Tldr:            "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla vitae elit libero, a pharetra augue. Donec id elit non mi porta gravida at eget metus. Nullam id dolor id nibh ultricies vehicula ut id elit.",
		Content:         "Lorem ipsum dolor sit MEHMET, consectetur adipiscing elit. Nulla vitae elit libero, a pharetra augue. Donec id elit non mi porta gravida at eget metus. Nullam id dolor id nibh ultricies vehicula ut id elit.",
		Url:             "blog/Title2",
		Tags:            []string{"tag1", "tag2"},
		Image:           "https://picsum.photos/200/300",
		Links:           []string{"link1", "link2"},
		Subheadings:     []string{"sub1", "sub2"},
		SubheadingLinks: []string{"sublink1", "sublink2"},
	})
	return blogs
}
