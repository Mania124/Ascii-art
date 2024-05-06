package main

import (
	"os"
	"strings"

	"ascii-art/asciiart"
)

func main() {
	fileName := asciiart.GetTheCorrectFile(os.Args)
	filePtr, err := os.ReadFile(fileName)
	asciiart.HandleError(err)
	if fileName == "thinkertoy.txt" {
		fileContents := strings.Split(string(filePtr), "\r\n")
		asciiMap := asciiart.MakeAsciiGraphicsMap(fileContents)
		asciiart.Output(os.Args[1], fileContents, asciiMap)
	} else {
		fileContents := strings.Split(string(filePtr), "\n")
		asciiMap := asciiart.MakeAsciiGraphicsMap(fileContents)
		asciiart.Output(os.Args[1], fileContents, asciiMap)
	}
}
