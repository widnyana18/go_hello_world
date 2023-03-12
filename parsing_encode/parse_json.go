package main

import (
	"encoding/json"
	"log"
	"net/url"
)

type User struct {
	Name string
	Age  int
}

func main() {
	//Parsing URL
	var urlString = "https://google.com/search?keyword=golang&page=2"
	link, err := url.Parse(urlString)

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("Parsing url: ", link)
	log.Println("Schema: ", link.Scheme)
	log.Println("host: ", link.Host)
	log.Println("path: ", link.Path)
	log.Println("Query String: ", link.Query())

	// Decode JSON to Struct
	var jsonString = `{"name": "Robert", "age": 24}`
	var jsonData = []byte(jsonString)

	var userStruct User
	err = json.Unmarshal(jsonData, &userStruct)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("struct: ", userStruct.Name)

	// Decode JSON to Map
	var userMap map[string]interface{}
	err = json.Unmarshal(jsonData, &userMap)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("map: ", userMap)

	// Decode JSON to Interface
	var userInter interface{}
	err = json.Unmarshal(jsonData, &userInter)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("interface: ", userInter)

	// Decode JSON to Array
	var userArray []User
	err = json.Unmarshal(jsonData, &userArray)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("array: ", userArray)

	// Encode struct to JSON
	var objUser = []User{{"Michael", 20}, {"Diana", 22}}

	jsonDt, err := json.Marshal(objUser)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println("json: ", string(jsonDt))
}
