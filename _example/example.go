package main

import (
	"fmt"
	"io/ioutil"

	svg "github.com/h2non/go-is-svg"
)

func main() {
	buf, err := ioutil.ReadFile("_example/example.svg")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	if svg.Is(buf) {
		fmt.Println("File is an SVG")
	} else {
		fmt.Println("File is NOT an SVG")
	}
}
