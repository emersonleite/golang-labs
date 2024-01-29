package countElementsOfAString

func CountElementsOfAString(anyString string) map[string]int16 {
	var count = map[string]int16{}
	for _, element := range anyString {
		key := string(element)
		count[key] += 1
	}
	return count
}
