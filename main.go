package main

import (
	"strings"
	"syscall/js"

	uhtml "tomdeneire.github.io/ulpia/uhtml"
	ujson "tomdeneire.github.io/ulpia/ujson"
	uxml "tomdeneire.github.io/ulpia/uxml"
)

// Handle API response from JS
func handleResponse(this js.Value, args []js.Value) interface{} {
	var source uhtml.Source
	source.Name = js.ValueOf(args[0]).String()
	source.Explain = js.ValueOf(args[1]).String()
	source.Format = js.ValueOf(args[2]).String()
	rawQuery := js.ValueOf(args[3]).String()
	response := []byte(js.ValueOf(args[4]).String())

	var result uhtml.Result

	if source.Format == "xml" {
		result, _ = uxml.ParseXML(response, source.Name)
	} else if source.Format == "json" {
		result, _ = ujson.ParseJSON(response, source.Name)
	}

	AddToResult(source, rawQuery, result)

	return nil
}

// Show result in "result" table
func AddToResult(source uhtml.Source, rawQuery string, result uhtml.Result) {
	table := js.Global().Get("document").Call("getElementById", "result")

	html := table.Get("innerHTML").String()

	html = html + "<tr><td><b>" +
		"<a href=\"" + source.Explain + "\" target=\"_blank\">" + source.Name + "</a></td><td>" +
		"<a href=\"" + rawQuery + "\" target=\"_blank\">raw</td><td><b>" +
		"identifiers" + "</td><td><b>" +
		"author" + "</td><td><b>" +
		"title" + "</td><td><b>" +
		"imprint" + "</td><td><b>" +
		"date" + "</td>" +
		"</tr>"

	for index, title := range result.Titles {
		identifiers := result.Identifiers[index]
		authors := result.Authors[index]
		imprints := result.Imprints[index]
		dates := result.Dates[index]
		title = title
		html = html +
			"<tr><td></td><td></td><td>" +
			identifiers + "</td><td>" +
			authors + "</td><td>" +
			title + "</td><td>" +
			imprints + "</td><td>" +
			dates + "</td>" +
			"</tr>"
	}

	html = strings.ReplaceAll(html, "\n", "<br>")

	table.Set("innerHTML", html)
}

// Register JS function callbacks
func registerCallbacks() {
	js.Global().Set("handleResponse", js.FuncOf(handleResponse))
}

// Main
func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
