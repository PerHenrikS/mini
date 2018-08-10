package main

import (
	"fmt"
	"mini/generator"
	"mini/manager"
	"os"
)

func main() {
	const usageHelp = `
commands: 
	mini gen  -- generates web page
	mini init -- initializes folder structure
	
usage:
	To create a webpage run 'mini init' inside 
	the directory you want the page to be created.
	
	Create posts with the format: 
	'postnumber-name-year.md' 
	
	Run mini gen to generate page 
	Serve folder
	`

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println(usageHelp)
		os.Exit(0)
	}

	switch args[0] {
	case "gen":
		mini := generator.New()
		mini.GeneratePage()
		fmt.Println("Page generated into ./webpage")
	case "init":
		manager.InitializeDir()
	default:
		fmt.Println(usageHelp)
	}
}
