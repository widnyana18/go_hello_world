package main

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/beevik/etree"
)

type Document struct {
	Debug  string
	Window struct {
		Title  string
		Name   string
		Width  int
		Height int
	}
	Image struct {
		Src       string
		Name      string
		HOffset   int
		VOffset   int
		Alignment string
	}
	Text struct {
		Data      string
		Size      int
		Style     string
		Name      string
		HOffset   int
		VOffset   int
		Alignment string
		OnMouseUp string
	}
}

const widgetJson = `{
"debug": "on",
"window": {
	"title": "Sample Konfabulator Widget",
	"name": "main_window",
	"width": 500,
	"height": 500
},
"image": { 
	"src": "Images/Sun.png",
	"name": "sun1",
	"hOffset": 250,
	"vOffset": 250,
	"alignment": "center"
},
"text": {
	"data": "Click Here",
	"size": 36,
	"style": "bold",
	"name": "text1",
	"hOffset": 250,
	"vOffset": 100,
	"alignment": "center",
	"onMouseUp": "sun1.opacity = (sun1.opacity / 100) * 90;"
}
}`

func main() {

	dataDoc := Document{}
	err := json.Unmarshal([]byte(widgetJson), &dataDoc)
	if err != nil {
		log.Fatal(err.Error())
	}

	var doc = etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	widget := doc.CreateElement("widget")

	widget.CreateElement("debug").SetText(dataDoc.Debug)

	window := widget.CreateElement("window")
	window.CreateElement("title").SetText(dataDoc.Window.Title)
	window.CreateAttr("name", dataDoc.Window.Name)
	window.CreateElement("width").SetText(strconv.Itoa(dataDoc.Window.Width))
	window.CreateElement("height").SetText(strconv.Itoa(dataDoc.Window.Height))

	image := widget.CreateElement("image")
	image.CreateAttr("src", dataDoc.Image.Src)
	image.CreateAttr("name", dataDoc.Image.Name)
	image.CreateElement("hOffset").SetText(strconv.Itoa(dataDoc.Image.HOffset))
	image.CreateElement("vOffset").SetText(strconv.Itoa(dataDoc.Image.VOffset))
	image.CreateElement("alignment").SetText(dataDoc.Image.Alignment)

	text := widget.CreateElement("text")
	text.CreateAttr("data", dataDoc.Text.Data)
	text.CreateElement("size").SetText(strconv.Itoa(dataDoc.Text.Size))
	text.CreateElement("style").SetText(dataDoc.Text.Style)
	text.CreateElement("name").SetText(dataDoc.Text.Name)
	text.CreateElement("hOffset").SetText(strconv.Itoa(dataDoc.Text.HOffset))
	text.CreateElement("vOffset").SetText(strconv.Itoa(dataDoc.Text.VOffset))
	text.CreateElement("alignment").SetText(dataDoc.Text.Alignment)
	text.CreateElement("onMouseUp").SetText(dataDoc.Text.OnMouseUp)

	doc.Indent(2)
	err = doc.WriteToFile("output.xml")
	if err != nil {
		log.Fatal(err)
	}
}
