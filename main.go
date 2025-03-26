package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/davidthesurreal/crane"
	"github.com/fatih/color"
)

func main() {
	arg, err := crane.GetArg(2)
	dir, err2 := crane.GetArg(3)

	if err != nil {
		fmt.Println("Welcome to Vial")
		return
	}

	if arg == "new" {
		if err2 != nil {
			color.HiRed("Project name must be included!")
			color.Cyan("Hint: Do `vial %v <project_name>` instead", arg)
			return
		}
		createProject(strings.Split(dir, ".")[0])
	} else if arg == "help" {
		fmt.Println("Use help to display this message.")
		fmt.Println("Use `new <project-name>` to create a new Flask project ")
	} else {
		color.Red("Invalid argument `%v`, Use `vial help` for help.", arg)
	}
}

func createProject(path string) {
	paths := []string{
		"/website/templates", "/website/static",
		"/website/routes",
	}

	files := []string{
		"/main.py", "/website/__init__.py",
		"/website/routes/views.py", "/website/routes/auth.py",
		"/website/templates/index.html",
		"/website/static/style.css",
		"/website/static/script.js",
	}

	color.Green("Creating project...")
	for _, dpath := range paths {
		err := os.MkdirAll(path+dpath, os.ModePerm)
		if err != nil {
			color.Red("%v", err)
		}
	}

	for _, file := range files {
		_, err := os.Create(path + file)
		if err != nil {
			color.Red("%v", err)
		}
	}

	fmt.Println(path, "created successfully!")
}
