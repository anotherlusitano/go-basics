package main

import (
	"fmt"
	"os"

	"files.io/data"
	"files.io/utils"
)

func cStuff() {
	price := 34.044819

	stringPrice := fmt.Sprintf("%.2f", price)

	fmt.Println(stringPrice)
}

func main() {
	cStuff()

	rootPath, _ := os.Getwd()

	filePath := rootPath + "/data/text.txt"

	c, err := utils.ReadTextFile(filePath)

	// fmt.Println("Current working directory:", rootPath)

	if err == nil {
		fmt.Println(c)
		newContent := fmt.Sprintf("Original: %v\nDouble Original: %v%v", c, c, c)

		utils.WriteToFile(filePath+".output.txt", newContent)
	} else {
		panic(err)
	}

	data.Test()
}
