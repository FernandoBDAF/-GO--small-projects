package main

import (
	"FileReader/fileutils"
	"fmt"
	"os"
)

func main() {
	// Getwd() -> get working directory
	wd, _ := os.Getwd()
	c, err := fileutils.ReadTextFile(wd + "/data/example.txt")
	if err != nil {
		println(err)
	} else {
		println(c)
	}
	price := 34.066000
	newContent := fmt.Sprintf("Original: %s\nDouble Original: %s / %s\nPrice: %.4v\n", c, c, c, price)
	errWriting := fileutils.WriteToFile(wd + "/data/example2.txt", newContent)
	if errWriting != nil {
		println(errWriting)
	}
}