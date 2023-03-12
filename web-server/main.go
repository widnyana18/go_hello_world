package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type person struct {
	ID   string
	Name string
	Age  int
}

var data = []person{
	person{"K01", "Joe", 20},
	person{"K02", "Say", 18},
	person{"K03", "Noie", 48},
	person{"K04", "Sia", 12},
}

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hello World")
	})

	http.HandleFunc("/index", index)

	http.HandleFunc("/golang", func(res http.ResponseWriter, req *http.Request) {
		var result, err = template.ParseFiles("template.html")

		if err != nil {
			fmt.Println(err.Error())
		}

		var data = map[string]string{
			"firstName": "Joe",
			"lastName":  "Doe",
		}

		result.Execute(res, data)
	})

	http.HandleFunc("/users", users)
	http.HandleFunc("/user", userById)

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hey Index")
}

func users(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	if req.Method == "POST" {
		result, err := json.Marshal(data)

		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Write(result)
		return
	}

	http.Error(res, "no clue", http.StatusBadRequest)
}

func userById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	if req.Method == "POST" {
		var result []byte
		var err error

		id := req.FormValue("id")

		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(res, err.Error(), http.StatusInternalServerError)
					return
				}

				res.Write(result)
				return
			}
		}

		http.Error(res, "user not found", http.StatusBadRequest)
		return
	}

	http.Error(res, "no clue", http.StatusBadRequest)
}
