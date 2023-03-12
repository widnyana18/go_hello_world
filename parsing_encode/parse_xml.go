package main

import (
	"encoding/json"
	"log"

	"github.com/beevik/etree"
)

type M map[string]interface{}

func main() {
	var doc = etree.NewDocument()
	if err := doc.ReadFromFile("data.xml"); err != nil {
		log.Fatal(err)
	}

	root := doc.SelectElement("breakfast_menu")
	data := make([]M, 0)

	for _, food := range root.FindElements("//food") {
		row := make(M)
		row["name"] = food.SelectElement("name").Text()
		row["price"] = food.SelectElement("price").Text()
		row["desc"] = food.SelectElement("description").Text()
		row["cal"] = food.SelectElement("calories").Text()

		data = append(data, row)
	}

	bts, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("result: ", string(bts))
}
