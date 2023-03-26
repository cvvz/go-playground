package books

import "strings"

type Book struct {
	Title  string
	Author string
	Pages  int
}

type Category string

const (
	CategoryNovel      Category = "novel"
	CategoryShortStory Category = "short story"
)

func (b Book) Category() Category {
	if b.Pages > 300 {
		return CategoryNovel
	}
	return CategoryShortStory
}

func (b Book) AuthorLastName() string {
	return strings.Split(b.Author, " ")[1]
}

func (b Book) AuthorFirstName() string {
	return strings.Split(b.Author, " ")[0]
}
