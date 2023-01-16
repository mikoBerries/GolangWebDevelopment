//How to including func , pipeline, and nested template in template golang
package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var indexTemplate *template.Template

type human struct {
	Name  string
	Ages  int
	Hobby string
	View  bool
}
type car struct {
	Name  string
	Brand string
	Year  int
}

func init() {
	// template.Must(template.ParseGlob("./Template_File/*"))

	//FuncMap => map[string]interface{}
	//map of string for "key" and empty interface{} for "value"

	// myFuncMap := template.FuncMap{
	// 	"UpperCase": strings.ToUpper,
	// 	"FT":        firstThree,
	// }
	// tmplt.Funcs(myFuncMap)

	myFuncMap := template.FuncMap{
		"UpperCase":  strings.ToUpper,
		"FT":         firstThree,
		"FormatTime": formatTime,
	}
	//template.New(string) to get *template.Template address
	indexTemplate = template.Must(template.New("index").Funcs(myFuncMap).ParseGlob("./Template_file/*.gohtml"))
}

func main() {
	//anonymous struct for single use
	myCommunity := struct {
		Humans []human
		Cars   []car
		Times  time.Time
		Score1 int
		Score2 int
	}{
		[]human{
			{"agustinus", 30, "slepping", true},
			{"budi", 10, "walking", false},
			{"don joe", 33, "runing", true},
		},
		[]car{
			{"Prius", "Toyota", 2001},
			{"NSX", "Honda", 1998},
			{"Lancer", "Subaru", 2001},
		},
		time.Now(),

		10,
		100,
	}

	if err := indexTemplate.ExecuteTemplate(os.Stdout, "tplWithFunc.gohtml", myCommunity); err != nil {
		log.Fatalln("err on ExecuteTemplate :", err)
	}

}

//get first 3 letter from string
func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

//format time stamp
func formatTime(t time.Time) string {
	// format time
	//01/02 03:04:05PM '06 -0700
	//01: Month
	//02: Day
	//06: Year
	//03: Hour
	//04: Minute
	//05: Second
	//07: location
	// 01/17 12:45:20AM '23 +0700 => 17 jan
	return t.Format("02/01/2006 03:04:05PM ")
}

func (h human) DoubleAges() int {
	return h.Ages * 2
}
func (h human) TripleAges(i int) int {
	return i * 3
}
