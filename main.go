package main
import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("usage: go run . <string> [banner]")
		return
	}
	banner := "standard"
	if len(os.Args) == 3 {
		banner = os.Args[2]
	}

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		fmt.Println("invalid banner")
		return
	}
	fileName := banner + ".txt"
	bannerMap := loadBanner(fileName)
	if bannerMap == nil {
		return
	}
	data := os.Args[1]
	fmt.Print(generateArt(data, bannerMap))
}