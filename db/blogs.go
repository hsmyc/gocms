package db

import "hsmyc/htmx/models"

func GetBlogs() []models.Blogmodel {
	var blogs []models.Blogmodel
	blogs = append(blogs, models.Blogmodel{
		Title:           "Title1",
		Subtitle:        "Subtitle",
		Author:          "Author",
		Time:            "02-01-2024",
		TldrHeader:      "TLDR:",
		Tldr:            "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla vitae elit libero, a pharetra augue. Donec id elit non mi porta gravida at eget metus. Nullam id dolor id nibh ultricies vehicula ut id elit.",
		Content:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ornare arcu dui vivamus arcu felis. Fringilla urna porttitor rhoncus dolor purus. Id ornare arcu odio ut sem nulla pharetra. Massa enim nec dui nunc mattis enim ut. Rhoncus aenean vel elit scelerisque mauris pellentesque. Amet luctus venenatis lectus magna fringilla urna porttitor rhoncus dolor. Massa vitae tortor condimentum lacinia quis vel eros donec ac. Odio ut enim blandit volutpat maecenas volutpat blandit. Morbi tristique senectus et netus et malesuada fames. Sit amet facilisis magna etiam tempor orci eu lobortis. Amet tellus cras adipiscing enim eu turpis egestas pretium. Parturient montes nascetur ridiculus mus mauris vitae. Mollis nunc sed id semper. Enim ut tellus elementum sagittis vitae et. Nibh sed pulvinar proin gravida hendrerit lectus a. Ac felis donec et odio pellentesque. Sed viverra tellus in hac habitasse platea. Morbi blandit cursus risus at. Sit amet nisl suscipit adipiscing bibendum est. Aliquam eleifend mi in nulla posuere sollicitudin aliquam. Nulla pellentesque dignissim enim sit amet venenatis urna cursus. Libero id faucibus nisl tincidunt eget nullam. Elit ut aliquam purus sit amet luctus venenatis lectus. Aliquam ut porttitor leo a diam sollicitudin tempor id eu. Amet mattis vulputate enim nulla aliquet porttitor lacus luctus. Faucibus interdum posuere lorem ipsum dolor sit amet. Sollicitudin ac orci phasellus egestas. Et ligula ullamcorper malesuada proin libero nunc consequat. Proin libero nunc consequat interdum varius sit. Neque sodales ut etiam sit amet. Cras sed felis eget velit aliquet sagittis id consectetur. Blandit libero volutpat sed cras ornare arcu dui vivamus arcu. Diam maecenas ultricies mi eget mauris. Viverra vitae congue eu consequat ac felis donec. Ante in nibh mauris cursus. Sed ullamcorper morbi tincidunt ornare massa eget egestas purus viverra. Mus mauris vitae ultricies leo integer malesuada. Feugiat sed lectus vestibulum mattis ullamcorper velit sed.Pulvinar pellentesque habitant morbi tristique senectus et netus et. Purus in massa tempor nec feugiat nisl pretium fusce id. Arcu dui vivamus arcu felis bibendum ut. Massa sapien faucibus et molestie ac feugiat sed lectus. Duis convallis convallis tellus id interdum velit. Feugiat pretium nibh ipsum consequat nisl vel pretium lectus. Malesuada fames ac turpis egestas integer eget. Arcu dictum varius duis at consectetur lorem donec. Sit amet aliquam id diam maecenas ultricies. Nulla porttitor massa id neque aliquam vestibulum morbi. Lectus sit amet est placerat in egestas erat imperdiet sed. Consequat mauris nunc congue nisi vitae suscipit tellus mauris a. Turpis egestas integer eget aliquet nibh. Ornare massa eget egestas purus viverra accumsan in. Interdum posuere lorem ipsum dolor sit amet.Nullam eget felis eget nunc lobortis. Ac ut consequat semper viverra. Nunc congue nisi vitae suscipit tellus. Massa sed elementum tempus egestas sed. Nulla malesuada pellentesque elit eget gravida. Luctus accumsan tortor posuere ac. Lorem ipsum dolor sit amet consectetur adipiscing. Ullamcorper velit sed ullamcorper morbi tincidunt. Sit amet mauris commodo quis imperdiet massa tincidunt nunc. Sit amet volutpat consequat mauris nunc congue nisi vitae suscipit. Neque laoreet suspendisse interdum consectetur libero id faucibus nisl. Sed cras ornare arcu dui vivamus arcu. Eu ultrices vitae auctor eu augue ut lectus arcu bibendum. Consectetur adipiscing elit ut aliquam. A lacus vestibulum sed arcu non odio euismod lacinia.",
		Url:             "blog/Title1",
		Tags:            []string{"tag1", "tag2"},
		Image:           "https://picsum.photos/600/500",
		Links:           []string{"link1", "link2"},
		Subheadings:     []string{"sub1", "sub2"},
		SubheadingLinks: []string{"sublink1", "sublink2"},
		Topics:          []string{"topic1", "topic2"},
		LikeCount:       "22",
		CommentCount:    "16",
		ShareCount:      "12",
	})
	blogs = append(blogs, models.Blogmodel{
		Title:           "Title2",
		Subtitle:        "Subtitle2",
		Author:          "Author2",
		Time:            "02-01-1991",
		TldrHeader:      "TLDR:",
		Tldr:            "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla vitae elit libero, a pharetra augue. Donec id elit non mi porta gravida at eget metus. Nullam id dolor id nibh ultricies vehicula ut id elit.",
		Content:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ornare arcu dui vivamus arcu felis. Fringilla urna porttitor rhoncus dolor purus. Id ornare arcu odio ut sem nulla pharetra. Massa enim nec dui nunc mattis enim ut. Rhoncus aenean vel elit scelerisque mauris pellentesque. Amet luctus venenatis lectus magna fringilla urna porttitor rhoncus dolor. Massa vitae tortor condimentum lacinia quis vel eros donec ac. Odio ut enim blandit volutpat maecenas volutpat blandit. Morbi tristique senectus et netus et malesuada fames. Sit amet facilisis magna etiam tempor orci eu lobortis. Amet tellus cras adipiscing enim eu turpis egestas pretium. Parturient montes nascetur ridiculus mus mauris vitae. Mollis nunc sed id semper. Enim ut tellus elementum sagittis vitae et. Nibh sed pulvinar proin gravida hendrerit lectus a. Ac felis donec et odio pellentesque. Sed viverra tellus in hac habitasse platea. Morbi blandit cursus risus at. Sit amet nisl suscipit adipiscing bibendum est. Aliquam eleifend mi in nulla posuere sollicitudin aliquam. Nulla pellentesque dignissim enim sit amet venenatis urna cursus. Libero id faucibus nisl tincidunt eget nullam. Elit ut aliquam purus sit amet luctus venenatis lectus. Aliquam ut porttitor leo a diam sollicitudin tempor id eu. Amet mattis vulputate enim nulla aliquet porttitor lacus luctus. Faucibus interdum posuere lorem ipsum dolor sit amet. Sollicitudin ac orci phasellus egestas. Et ligula ullamcorper malesuada proin libero nunc consequat. Proin libero nunc consequat interdum varius sit. Neque sodales ut etiam sit amet. Cras sed felis eget velit aliquet sagittis id consectetur. Blandit libero volutpat sed cras ornare arcu dui vivamus arcu. Diam maecenas ultricies mi eget mauris. Viverra vitae congue eu consequat ac felis donec. Ante in nibh mauris cursus. Sed ullamcorper morbi tincidunt ornare massa eget egestas purus viverra. Mus mauris vitae ultricies leo integer malesuada. Feugiat sed lectus vestibulum mattis ullamcorper velit sed.Pulvinar pellentesque habitant morbi tristique senectus et netus et. Purus in massa tempor nec feugiat nisl pretium fusce id. Arcu dui vivamus arcu felis bibendum ut. Massa sapien faucibus et molestie ac feugiat sed lectus. Duis convallis convallis tellus id interdum velit. Feugiat pretium nibh ipsum consequat nisl vel pretium lectus. Malesuada fames ac turpis egestas integer eget. Arcu dictum varius duis at consectetur lorem donec. Sit amet aliquam id diam maecenas ultricies. Nulla porttitor massa id neque aliquam vestibulum morbi. Lectus sit amet est placerat in egestas erat imperdiet sed. Consequat mauris nunc congue nisi vitae suscipit tellus mauris a. Turpis egestas integer eget aliquet nibh. Ornare massa eget egestas purus viverra accumsan in. Interdum posuere lorem ipsum dolor sit amet.Nullam eget felis eget nunc lobortis. Ac ut consequat semper viverra. Nunc congue nisi vitae suscipit tellus. Massa sed elementum tempus egestas sed. Nulla malesuada pellentesque elit eget gravida. Luctus accumsan tortor posuere ac. Lorem ipsum dolor sit amet consectetur adipiscing. Ullamcorper velit sed ullamcorper morbi tincidunt. Sit amet mauris commodo quis imperdiet massa tincidunt nunc. Sit amet volutpat consequat mauris nunc congue nisi vitae suscipit. Neque laoreet suspendisse interdum consectetur libero id faucibus nisl. Sed cras ornare arcu dui vivamus arcu. Eu ultrices vitae auctor eu augue ut lectus arcu bibendum. Consectetur adipiscing elit ut aliquam. A lacus vestibulum sed arcu non odio euismod lacinia.",
		Url:             "blog/Title2",
		Tags:            []string{"tag1", "tag2"},
		Image:           "https://picsum.photos/600/500",
		Links:           []string{"link1", "link2"},
		Subheadings:     []string{"sub1", "sub2"},
		SubheadingLinks: []string{"sublink1", "sublink2"},
		Topics:          []string{"topic1", "topic2"},
		LikeCount:       "22",
		CommentCount:    "16",
		ShareCount:      "12",
	})
	return blogs
}
