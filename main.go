package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	uxml "./lib/xml"

	"github.com/clbanning/mxj"
)

// Handle XML response from JS
func handleXML(this js.Value, args []js.Value) interface{} {
	identifier := js.ValueOf(args[0]).String()
	response := []byte(js.ValueOf(args[1]).String())
	mv, err := mxj.NewMapXml(response)
	if err != nil {
		return fmt.Errorf("cannot parse XML: %v", err)
	}
	result, err := json.Marshal(mv)
	if err != nil {
		return fmt.Errorf("invalid JSON: %v", err)
	}

	// parse results according to identifier (XML structure not always the same)
	_ = uxml.ParseXML(response)
	addToResult(identifier, string(result))
	return nil
}

// Register JS function callbacks
func registerCallbacks() {
	js.Global().Set("handleXML", js.FuncOf(handleXML))
}

// Show result in "result" table
func addToResult(identifier string, result string) {
	table := js.Global().Get("document").Call("getElementById", "result")
	result = table.Get("innerHTML").String() + "<tr><td>" + identifier + "</td><td>" + result + "</td></tr>"
	table.Set("innerHTML", result)
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
