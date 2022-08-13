package main

import (
	"syscall/js"

	uhtml "tomdeneire.github.io/ulpia/uhtml"
	uxml "tomdeneire.github.io/ulpia/uxml"
)

// Handle XML response from JS
func handleXML(this js.Value, args []js.Value) interface{} {
	source := js.ValueOf(args[0]).String()
	response := []byte(js.ValueOf(args[1]).String())

	result, _ := uxml.ParseXML(response, source)
	uhtml.AddToResult(source, result)

	return nil
}

// Register JS function callbacks
func registerCallbacks() {
	js.Global().Set("handleXML", js.FuncOf(handleXML))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
