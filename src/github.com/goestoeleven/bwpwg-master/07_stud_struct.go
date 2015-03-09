package main

import (
	"log"
	"os"
	"text/template"
)

type Student struct {
	//exported field since it begins
	//with a capital letter
	Name string
}

func main() {
	//define an instance
	s := Student{"Satish"}

	// STEP 1: create a new template
	// STEP 2: parse the string into the template
	// STEP 3: execute the template

	// STEP 1: create a new template with some name
	tmpl := template.New("test")

	// STEP 2: parse the string into the template
	tmpl, err := tmpl.Parse("Hello {{.Name}}!")
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	// STEP 3: execute the template
	//merge template 'tmpl' with content of 's'
	err1 := tmpl.Execute(os.Stdout, s)
	if err1 != nil {
		log.Fatal("Execute: ", err1)
		return
	}
}
