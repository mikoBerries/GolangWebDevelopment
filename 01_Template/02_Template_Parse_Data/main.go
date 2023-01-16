//execute template and trow data to template
package main

import (
	"log"
	"os"
	"text/template"
)

var myTemplate *template.Template
var mySliceData []string
var myMapData map[string]string

type human struct {
	Name  string
	Ages  int
	Hobby string
}
type car struct {
	Name  string
	Brand string
	Year  int
}

// type community struct {
// 	Humans []human
// 	Cars   []car
// }

var tempHuman human

// var myCommunity community

func init() {
	//parsing all template
	myTemplate = template.Must(template.ParseGlob("./TemplateFile/*"))
	mySliceData = []string{"data one", "data two", "data three"}
	myMapData = map[string]string{
		"Name":  "Budi",
		"Ages":  "29",
		"hobby": "swimming",
		"pet":   "cat",
	}
	tempHuman = human{
		"agustinus",
		30,
		"sleeping",
	}
	// myCommunity = community{
	// 	[]human{
	// 		{"agustinus", 30, "slepping"},
	// 		{"budi", 10, "walking"},
	// 		{"don joe", 33, "runing"},
	// 	},
	// 	[]car{
	// 		{"Prius", "Toyota", 2001},
	// 		{"NSX", "Honda", 1998},
	// 		{"Lancer", "Subaru", 2001},
	// 	},
	// }
}

func main() {
	err := myTemplate.ExecuteTemplate(os.Stdout, "tplSliceString.gohtml", mySliceData)
	if err != nil {
		log.Fatalln("err when execute Template tplSliceString: ", err)
	}

	err = myTemplate.ExecuteTemplate(os.Stdout, "tplMapString.gohtml", myMapData)
	if err != nil {
		log.Fatalln("err when execute Template tplMapString: ", err)
	}

	err = myTemplate.ExecuteTemplate(os.Stdout, "tplStruct.gohtml", tempHuman)
	if err != nil {
		log.Fatalln("err when execute Template tplStruct: ", err)
	}
	//anonymous struct for single use
	myCommunity := struct {
		Humans []human
		Cars   []car
	}{
		[]human{
			{"agustinus", 30, "slepping"},
			{"budi", 10, "walking"},
			{"don joe", 33, "runing"},
		},
		[]car{
			{"Prius", "Toyota", 2001},
			{"NSX", "Honda", 1998},
			{"Lancer", "Subaru", 2001},
		},
	}
	err = myTemplate.ExecuteTemplate(os.Stdout, "tplStructSliceStruct.gohtml", myCommunity)
	if err != nil {
		log.Fatalln("err when execute Template tplStruct: ", err)
	}
}
