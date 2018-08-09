package generator

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"strings"

	blackfriday "gopkg.in/russross/blackfriday.v2"
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
	Content template.HTML
}

//NewPost : creates a Post structure from a filepath
func NewPost(path string, filename string) (Post, error) {
	postInfo := strings.Split(filename, "-")
	if len(postInfo) != 3 {
		return Post{}, fmt.Errorf("File name of post must be in format '#-postname-year.md' \nCheck your posts directory")
	}

	var newPost Post
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return Post{}, fmt.Errorf("Could not read file from path: " + path)
	}

	input := []byte(fileContent)

	output := blackfriday.Run(input)
	newPost = Post{
		Title:   postInfo[1],
		Number:  postInfo[0],
		Year:    strings.TrimSuffix(postInfo[2], ".md"),
		Content: template.HTML(output)}
	return newPost, nil
}

//NewPage : helper function to create a webpage structure
func NewPage(title string, author string, authoremail string, posts []Post) Page {
	return Page{Title: title, Author: author, Posts: posts, AuthorEmail: authoremail}
}
