package main

import (
	"fmt"
	"os"

	"files.io/utils"
)

func main() {
	rootPath, _ := os.Getwd()

	// fmt.Println("Current working directory:", rootPath)

	content, err := utils.ReadTextFile(rootPath + "/data/text.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(content)
}
