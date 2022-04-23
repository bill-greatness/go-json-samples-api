package posts

import "time"

var TempPost = []Post{
	{
		UserID:  1,
		Title:   "Getting Started in Life",
		Content: "This is a very long text content",
		Date:    time.Now(),
	},
	{
		UserID:  2,
		Title:   "Getting Started in Life 2",
		Content: "This is a very long text content",
		Date:    time.Date(2021, time.April, 4, 4, 5, 3, 5, time.Local),
	},
	{
		UserID:  3,
		Title:   "Getting Started in Life 3",
		Content: "This is a very long text content",
		Date:    time.Date(2021, time.April, 4, 4, 5, 13, 5, time.Local),
	},
}

func GetData(total int) []Post {

	return TempPost
}
