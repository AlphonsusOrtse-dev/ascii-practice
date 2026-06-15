package main

import "strings"

func rederLines(lines string, bannerMap map[rune][]string)string{
	var sb strings.Builder
	if lines == "" {
		for i := 0; i < 8; i++{
			sb.WriteString("\n")
		}
		return sb.String()
	}
	for row := 0; row < 8; row++ {
		for _, ch := range lines {
			sb.WriteString(bannerMap[ch][row])
		}
		sb.WriteString("\n")
	}
	return sb.String()
}