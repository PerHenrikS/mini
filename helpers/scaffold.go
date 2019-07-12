package helpers

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const welcomePost = `
# Welcome 

This is a post, follow the file naming conventions and all will be good.
`

//InitDir : initializes directory structure with layout directory and contents
func InitDir() {
	CreateDir("./assets")
	if CreateDir("./posts") {
		if err := CreateAndWrite("./posts/1-Welcome!-2018.md", welcomePost); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("======Ready======")
		fmt.Println("Open up './posts/1-Welcome!-2018.md' and start exploring!")
	} else {
		fmt.Println("posts directory already exists")
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
