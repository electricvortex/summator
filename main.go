package main

import (
	"fmt"
	"github.com/otiai10/gosseract"
)

// It is working using tesseract

func main() {
	client := gosseract.NewClient()
	err := client.SetLanguage("rus", "eng")
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()
	err = client.SetImage("/Users/sattar/Downloads/634802986_2730330.jpg")
	if err != nil {
		fmt.Println(err)
	}
	text, _ := client.Text()
	fmt.Println(text)
}
