package uhtml

import (
	"strings"
	"syscall/js"

	uxml "tomdeneire.github.io/ulpia/uxml"
)

// Show result in "result" table
func AddToResult(source string, result uxml.Result) {
	table := js.Global().Get("document").Call("getElementById", "result")

	html := table.Get("innerHTML").String()

	html = html + "<tr><td><b>" +
		source + "</td><td><b>" +
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
			"<tr><td></td><td>" +
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
