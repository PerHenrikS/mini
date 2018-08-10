package generator

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"mini/helpers"
	"os"
	"strings"

	"github.com/tkanos/gonfig"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

//Configuration : contains paths from json
type Configuration struct {
	PageTitle         string
	AuthorEmail       string
	PageAuthor        string
	PagesPath         string
	TemplateIndexPath string
	TemplatePostPath  string
	TemplateStylePath string
	WebpagePath       string
	PostsPath         string
	StylePath         string
}

//Generator : struct to contain information needed to page generation
type Generator struct {
	config Configuration
}

//New : creates a page generator
func New() Generator {
	generator := Generator{}
	err := gonfig.GetConf("./conf/config.json", &generator.config)
	helpers.Check(err)

	return generator
}

//GeneratePage : Creates a folder with webpage
func (g *Generator) GeneratePage() {
	posts := []Post{}

	files, err := ioutil.ReadDir(g.config.PagesPath)
	helpers.Check(err)

	helpers.CreateDir(g.config.WebpagePath)
	helpers.CreateDir(g.config.StylePath)
	helpers.CopyFile(g.config.TemplateStylePath, g.config.StylePath+"/style.css")

	os.RemoveAll(g.config.PostsPath) //To not get duplicates, inefficient but works for now
	helpers.CreateDir(g.config.PostsPath)

	for _, f := range files {
		var path = g.config.PagesPath + "/" + f.Name()
		newPost, err := g.newPost(path, f.Name())
		helpers.Check(err)
		posts = append(posts, newPost)

		templateContent, err := ioutil.ReadFile(g.config.TemplatePostPath)
		helpers.Check(err)

		t, err := template.New("Post").Parse(string(templateContent))
		helpers.Check(err)

		file, err := os.Create("./webpage/posts/" + newPost.Number + "-" + newPost.Title + ".html")
		helpers.Check(err)
		defer file.Close()

		err = t.Execute(file, newPost)
		helpers.Check(err)
	}

	webpage := NewPage(g.config.PageTitle, g.config.PageAuthor, g.config.AuthorEmail, posts)

	fileContent, err := ioutil.ReadFile(g.config.TemplateIndexPath)
	helpers.Check(err)

	t, err := template.New("Homepage").Parse(string(fileContent))
	helpers.Check(err)

	file, err := os.Create("./webpage/index.html")
	helpers.Check(err)
	defer file.Close()

	err = t.Execute(file, webpage)
	helpers.Check(err)
}

func (g *Generator) newPost(path string, filename string) (Post, error) {
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
		Author:  g.config.PageAuthor,
		Year:    strings.TrimSuffix(postInfo[2], ".md"),
		Content: template.HTML(output)}
	return newPost, nil
}
