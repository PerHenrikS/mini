package generator

import (
	"html/template"
)

//Page : webpage structure
type Page struct {
	Title       string
	Author      string
	AuthorEmail string
	Posts       []Post
}

//Post : structure that contains a post page
type Post struct {
	Title   string
	Number  string
	Year    string
	Author  string
	Content template.HTML
}

//NewPage : helper function to create a webpage structure
func NewPage(title string, author string, authoremail string, posts []Post) Page {
	return Page{Title: title, Author: author, Posts: posts, AuthorEmail: authoremail}
}
