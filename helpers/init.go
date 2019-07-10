package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/AlecAivazis/survey.v1"
)

//Struct to generate config.json
type genConfig struct {
	PageTitle         string `json:"PageTitle"`
	PageAuthor        string `json:"PageAuthor"`
	AuthorEmail       string `json:"AuthorEmail"`
	PagesPath         string `json:"PagesPath"`
	TemplateIndexPath string `json:"TemplateIndexPath"`
	TemplatePostPath  string `json:"TemplatePostPath"`
	TemplateStylePath string `json:"TemplateStylePath"`
	WebpagePath       string `json:"WebpagePath"`
	PostsPath         string `json:"PostsPath"`
	AssetsPath        string `json:"AssetsPath"`
	StylePath         string `json:"StylePath"`
	Port              string `json:"Port"`
}

/*
These are constants for now, I don't want the user to change these
until a better directory system is implemented
*/
const pagesPath = "./posts"
const templateIndexPath = "./layout/index.html"
const templatePostPath = "./layout/post.html"
const templateStylePath = "./layout/css/main.css"
const webpagePath = "./webpage"
const assetsPath = "./webpage/assets"
const postsPath = "./webpage/posts"
const stylePath = "./webpage/css"
const defaultPort = ":8000"

//Questions for the init survey
var qs = []*survey.Question{
	{
		Name: "Author",
		Prompt: &survey.Input{
			Message: "Author name:",
		},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "Pagename",
		Prompt: &survey.Input{
			Message: "Webpage title:",
		},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "Email",
		Prompt: &survey.Input{
			Message: "Author email: ",
		},
		Transform: survey.Title,
	},
}

//InitConf : initializes directory for generator
func InitConf() {
	answers := struct {
		Author   string
		Pagename string
		Email    string
	}{}

	err := survey.Ask(qs, &answers)
	Check(err)

	//TODO: Some message if it already exists?
	CreateDir("conf")

	configuration := genConfig{
		PageTitle:         answers.Pagename,
		PageAuthor:        answers.Author,
		AuthorEmail:       answers.Email,
		PagesPath:         pagesPath,
		TemplateIndexPath: templateIndexPath,
		TemplatePostPath:  templatePostPath,
		TemplateStylePath: templateStylePath,
		WebpagePath:       webpagePath,
		PostsPath:         postsPath,
		AssetsPath:        assetsPath,
		StylePath:         stylePath,
		Port:              defaultPort,
	}

	jsonContent, err := json.MarshalIndent(&configuration, "", "\t\t")
	Check(err)

	err = ioutil.WriteFile("conf/config.json", jsonContent, 0644)
	Check(err)

	fmt.Printf(`
==================
Config initialized
==================

Author: %s
Pagename: %s
Email: %s

`,
		answers.Author,
		answers.Pagename,
		answers.Email,
	)

	/*
		InitDir is not supposed to do anything if layout already exists.
		I don't want to override the users layout if a custom one
		is made.
		It is only made for initialization.
	*/
	InitDir()
}
