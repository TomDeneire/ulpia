package ujson

import (
	"encoding/json"

	uhtml "tomdeneire.github.io/ulpia/uhtml"
)

// Parse JSON responses according to specification
func ParseJSON(data []byte, source string) (uhtml.Result, error) {
	var result uhtml.Result

	if source == "europeana" {
		var response Europeana
		err := json.Unmarshal(data, &response)
		if err != nil {
			return result, err
		}
		for _, record := range response.Items {
			result.Titles = append(result.Titles,
				loopConCat(record.Title))
			result.Identifiers = append(result.Identifiers,
				record.ID)

			result.Dates = append(result.Dates,
				loopConCat(record.Year))

			result.Authors = append(result.Authors,
				loopConCat(record.DcCreator))

			result.Imprints = append(result.Imprints,
				"")
		}
		return result, nil
	}

	return result, nil
}

func loopConCat(data []string) string {
	var result string
	for _, item := range data {
		result += "\n" + item
	}
	return result
}
