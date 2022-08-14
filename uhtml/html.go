package uhtml

import (
	"strings"
	"syscall/js"
)

type Source struct {
	Name    string
	Type    string
	Explain string
}

type Result struct {
	Titles      []string
	Dates       []string
	Authors     []string
	Identifiers []string
	Imprints    []string
}

// Show result in "result" table
func AddToResult(source Source, rawQuery string, result Result) {
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
