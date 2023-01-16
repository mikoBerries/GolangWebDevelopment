//Template golang
package main

import (
	"log"
	"os"
	"text/template"
)

var tmplate *template.Template

func init() {
	// template.Must is error checking when parse all template and will panic if some file returning erorr(checker)
	tmplate = template.Must(template.ParseFiles("./TemplateFile/tpl.gohtml"))
}

func main() {
	// tmplate, err := template.ParseFiles("./TemplateFile/tpl.gohtml")
	// if err != nil {
	// 	log.Fatalf("err (template.ParseFiles): %v", err)
	// }

	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalf("err on (os.Create): %v", err)
	}
	defer newFile.Close()
	err = tmplate.Execute(os.Stdout, `self minding;self control`)

	if err != nil {
		log.Fatalf("err func (tmplate.Execute): %v", err)
	}
	//parse golb for big data
	tmplate, err = template.ParseGlob("./TemplateFile/*")
	if err != nil {
		log.Fatalf("err (template.ParseGlob): %v", err)
	}
	//tmplate value has 3 template and can execute secify template call by name of file
	tmplate.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalf("err (template.ExecuteTemplate): %v", err)
	}

	tmplate.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalf("err (template.ExecuteTemplate): %v", err)
	}
}
