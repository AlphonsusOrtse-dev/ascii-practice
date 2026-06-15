package main

import (
	"fmt"
	"os"
	"strings"
)

func loadBanner(s string)map[rune][]string{
	read, err := os.ReadFile(s)
	if err != nil {
		fmt.Println("error reading file", err)
		return nil
	}
	data := string(read)
	data = strings.ReplaceAll(data, "\r\n", "\n")
	if len(strings.TrimSpace(data))== 0 {
		fmt.Println("bamner file is empty")
		return nil
	}
	blocks := strings.Split(data, "\n\n")
	result := map[rune][]string{}
	for index, block := range blocks{
		if index >= 95 {
			break
		}
		ch := ' ' + rune(index)
		lines := strings.Split(block, "\n")
		if len(lines) < 8 {
			fmt.Println("incomplete baner file")
			return nil
		}
		result[ch] = lines[:8]
	}
	return  result
	

}