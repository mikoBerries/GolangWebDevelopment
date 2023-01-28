/*
	basic usage json
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//unmarshaling json using tag
type city struct {
	Bali       string  `json:"Postal"`
	Kauai      float64 `json:"Latitude"`
	Maui       float64 `json:"Longitude"`
	Java       string  `json:"Address"`
	NewZealand string  `json:"City"`
	Skye       string  `json:"State"`
	Oahu       string  `json:"Zip"`
	Hawaii     string  `json:"Country"`
}

type cities []city

type model struct {
	State    bool
	Pictures []string
}

func main() {
	var data cities

	rcvd := `[{"Postal":"zip","Latitude":37.7668,"Longitude":-122.3959,"Address":"","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},{"Postal":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"","City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`
	//unmarshal incoming json data
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)
	fmt.Println(data[1].Kauai)

	m := model{
		State: true,
		Pictures: []string{
			"one.jpg",
			"two.jpg",
			"three.jpg",
		},
	}

	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}

	os.Stdout.Write(bs)
}
