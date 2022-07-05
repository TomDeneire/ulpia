package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall/js"
)

// Handle JS input
func handleInput(this js.Value, input []js.Value) interface{} {
	author := js.ValueOf(input[0]).String()
	title := js.ValueOf(input[1]).String()
	query := author + "%20AND%20" + title
	result, err := getUnicat(query)
	if err != nil {
		panic(err)
	}
	setResult(result)
	return nil
}

// Register JS function callbacks
func registerCallbacks() {
	js.Global().Set("handleInput", js.FuncOf(handleInput))
}

func getUnicat(query string) (string, error) {
	url := "https://www.unicat.be/sru?version=1.1&operation=searchRetrieve&query=" + query + "&recordSchema=dc"
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// Manipulate DOM to show output
func setResult(result string) {
	output := js.Global().Get("document").Call("getElementById", "result")
	output.Set("innerText", result)
}

func main() {
	c := make(chan struct{}, 0)
	fmt.Println("WASM Go Initialized")
	registerCallbacks()
	<-c
}
