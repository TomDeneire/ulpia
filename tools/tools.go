package tools

func LoopConCat(data []string) string {
	var result string
	for _, item := range data {
		result += "\n" + item
	}
	return result
}
