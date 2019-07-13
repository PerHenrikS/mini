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
	PageTitle   string
	AuthorEmail string
	PageAuthor  string
	PagesPath   string
	ThemeName   string
	Port        string
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
	files, err := ioutil.ReadDir("./posts")
	helpers.Check(err)

	/*
		FIXME:
			Add the directories based on the layout
			directory structure as the webpage
			will be generated from the existing template.
	*/
	helpers.CreateDir("./webpage")

	//helpers.CreateDir(g.config.StylePath)	FIXME: Add these iteratively based on theme structure
	helpers.CreateDir("./assets")

	// FIXME: Copy the css file / files from ./layout/themename --> webpage/...
	//helpers.CopyFile(g.config.TemplateStylePath, g.config.StylePath+"/style.css")
	helpers.CreateDir("./webpage/css")
	helpers.CopyFile("./layout/"+g.config.ThemeName+"/css/main.css", "./webpage/css/style.css")

	// This refers to the compiled posts
	os.RemoveAll("./webpage/posts") //To not get duplicates, inefficient but works for now
	helpers.CreateDir("./webpage/posts")

	posts := g.generatePosts(files)

	webpage := NewPage(g.config.PageTitle, g.config.PageAuthor, g.config.AuthorEmail, posts)

	/*
		FIXME:
			Read from layout/theme/index.html
	*/
	fileContent, err := ioutil.ReadFile("./layout/" + g.config.ThemeName + "/index.html") // FIXME: Initial test with default layout
	helpers.Check(err)

	t, err := template.New("Homepage").Parse(string(fileContent))
	helpers.Check(err)

	file, err := os.Create("./webpage/index.html")
	helpers.Check(err)
	defer file.Close()

	err = t.Execute(file, webpage)
	helpers.Check(err)
}

/*
Creates a Post structure from a .md file
*/
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

/*
Creates Post structures out of .md files defined in posts directory
and generates the post.html pages
*/
func (g *Generator) generatePosts(files []os.FileInfo) []Post {
	posts := []Post{}
	for _, f := range files {
		var path = "./posts/" + f.Name()
		newPost, err := g.newPost(path, f.Name())
		helpers.Check(err)
		posts = append(posts, newPost)

		// FIXME: Take from layout/<theme>/post
		templateContent, err := ioutil.ReadFile("./layout/" + g.config.ThemeName + "/post.html")
		helpers.Check(err)

		t, err := template.New("Post").Parse(string(templateContent))
		helpers.Check(err)

		file, err := os.Create("./webpage/posts/" + newPost.Number + "-" + newPost.Title + ".html")
		helpers.Check(err)
		defer file.Close()

		err = t.Execute(file, newPost)
		helpers.Check(err)
	}

	return posts
}
