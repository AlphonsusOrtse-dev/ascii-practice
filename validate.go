package main


func validateInput(input string, bannerMap map[rune][]string)bool {
	for _, ch := range input{
		_, ok := bannerMap[ch]
		if !ok {
			return false
		}
	}
	return  true
}