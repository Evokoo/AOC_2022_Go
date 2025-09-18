package tools

import (
	"os"
	"regexp"
)

func ReadFile(title string) string {
	data, err := os.ReadFile(title)
	if err != nil {
		panic("Error reading file")
	}
	return string(data)
}

func QuickMatch(str, pattern string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindAllString(str, -1)
}
