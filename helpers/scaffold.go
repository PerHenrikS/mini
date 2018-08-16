package helpers

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/*
This does not feel too elegant - how else would this be done?
*/
const index = `
<!doctype html>
<html>
  <head>
    <meta charset="urf-8">
    <title>{{.Title}}</title>
    <link href="https://fonts.googleapis.com/css?family=Karla" rel="stylesheet">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="/css/style.css">
  </head>
  <body>
    <div class="page-container">
      <div class="page-header">
        <h1 class="header-title">{{.Title}}</h1>
        <small>{{.Author}}</small>
      </div>
      <div class="posts-list">
        {{range .Posts}}
        <a class="post-display" href="/posts/{{.Number}}-{{.Title}}.html">
          <div class="post-display">
            <h4>{{.Title}}</h4>
          </div>
        </a>
        {{end}}
    </div>
  </body>
</html>	
`

const post = `
<!doctype html>
<html>
  <head>
    <meta charset="urf-8">
    <title>{{.Title}}</title>
    <link href="https://fonts.googleapis.com/css?family=Karla" rel="stylesheet">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="/css/style.css">
  </head>
  <body>
    <div class="page-header">
      <a href="/">Home</a>
      <h1 class="header-title">{{.Title}} - {{.Year}}</h1>
      <small>
        {{.Author}}
      </small>
    </div>
    <div class="post-container">
      {{.Content}}
    </div>
  </body>
</html>
`

const styling = `
body{
  padding: 0; 
  margin: 0; 
  font-family: 'Karla', sans-serif;
}
.page-container {
  display: flex; 
  flex-direction: column;
}
/*HEADER*/
.page-header {
  display: flex; 
  flex-direction: column; 
  align-items: center;
  justify-content: center; 
  background-color: #1d1d1d;
  color: white;
}
.page-header a {
  text-decoration: none;
  color: white; 
}
.header-title {
  margin: 0; 
}
/*HEADER END*/
/*MAIN CONTENT*/
.posts-list {
  display: flex; 
  justify-content: center;
  align-items: center;
  flex-direction: column; 
}
.post-display{
  display: flex; 
  justify-content: center; 
  align-items: center;
  height: 4em;
  width: 50vw;
  color: black;   
  text-decoration: none;  
  transition-property: width; 
  transition-duration: 0.1s; 
  margin-bottom: 1px; 
}
.post-display:hover{
  width: 40vw;  
  border-bottom: 1px solid black; 
  margin-bottom: 0px; 
}
/*MAIN END*/
/*POST*/
.post-container {
  display: flex; 
  flex-direction: column;
  padding-top: 3em; 
  padding-left: 20%; 
  padding-right: 20%;
  padding-bottom: 5em;
  justify-content: center;
  text-align: justify;
}
/*END POST*/
`

const welcomePost = `
# Welcome 

This is a post, follow the file naming conventions and all will be good.
`

//InitDir : initializes directory structure with layout directory and contents
func InitDir() {
	//If directory is created
	if CreateDir("./layout") {
		CreateDir("./layout/css")
		urls := [...]string{
			"./layout/index.html",
			"./layout/post.html",
			"./layout/css/main.css",
		}
		/*Should the key be more explicit? ie. the full pathname*/
		urlContents := map[string]string{
			urls[0]: index,
			urls[1]: post,
			urls[2]: styling,
		}

		for _, url := range urls {
			if err := CreateAndWrite(url, urlContents[url]); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		if CreateDir("./posts") {
			if err := CreateAndWrite("./posts/1-Welcome!-2018.md", welcomePost); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Println("posts directory already exists")
		}
	} else {
		fmt.Println("layout directory already exists.")
	}
}

//CreateAndWrite : creates file and writes content to it
func CreateAndWrite(path string, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, strings.NewReader(content))
	if err != nil {
		return err
	}

	return nil
}
