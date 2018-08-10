package helpers

import (
	"fmt"
	"io/ioutil"
	"os"
)

//CreateDir : creates a directory in desired path
func CreateDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0700) //Read Write Execute permissions TODO: Maybe change this to something more suitable?
	}
}

//CopyFile : copies a file to another
func CopyFile(src string, dst string) {
	input, err := ioutil.ReadFile(src)
	Check(err)

	err = ioutil.WriteFile(dst, input, 0644)
	Check(err)
}

//Check : checks error
func Check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
