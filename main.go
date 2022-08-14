package main

import (
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

	uhtml.AddToResult(source, rawQuery, result)

	return nil
}

// Register JS function callbacks
func registerCallbacks() {
	js.Global().Set("handleResponse", js.FuncOf(handleResponse))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
