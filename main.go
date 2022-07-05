package main

import (
	"fmt"
	"syscall/js"
)

// Handle JS input
func handle(this js.Value, input []js.Value) interface{} {
	val1 := js.ValueOf(input[0]).String()
	val2 := js.ValueOf(input[1]).String()
	result := val1 + val2
	setResult(result)
	return nil
}

// Register function callbacks
func registerCallbacks() {
	js.Global().Set("handle", js.FuncOf(handle))
}

// Manipualte DOM to show output
func setResult(result string) {
	output := js.Global().Get("document").Call("getElementById", "output")
	output.Set("value", result)
}

func main() {
	c := make(chan struct{}, 0)
	fmt.Println("WASM Go Initialized")
	registerCallbacks()
	<-c
}
