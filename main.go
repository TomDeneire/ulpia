package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/clbanning/mxj"
)

// Handle XML response from JS
func handleXML(this js.Value, table []js.Value) interface{} {
	response := []byte(js.ValueOf(table[0]).String())
	mv, err := mxj.NewMapXml(response)
	if err != nil {
		return fmt.Errorf("cannot parse XML: %v", err)
	}
	result, err := json.Marshal(mv)
	if err != nil {
		return fmt.Errorf("invalid JSON: %v", err)
	}

	// parse result in html table
	addToResult(string(result))
	return nil
}

// Register JS function callbacks
func registerCallbacks() {
	js.Global().Set("handleXML", js.FuncOf(handleXML))
}

// Show result in "result" table
func addToResult(result string) {
	table := js.Global().Get("document").Call("getElementById", "result")
	result = table.Get("innerText").String() + result
	table.Set("innerText", result)
}

func main() {
	c := make(chan struct{}, 0)
	fmt.Println("WASM Go Initialized")
	registerCallbacks()
	<-c
}
