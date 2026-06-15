package main


import "strings"

func generateArt(input string, bannerMap map[rune][]string)string{
	lines := splitInput(input)

	for _, line := range lines {
		if !validateInput(line, bannerMap){
			return "invalid character detected\n"
		}
	}
	var result strings.Builder
	for _, line := range lines{
		result.WriteString(rederLines(line, bannerMap))
	}
	return result.String()
}