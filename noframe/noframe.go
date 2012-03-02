package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
)

func HandleImage(input string) (err error) {
	var f *os.File
	if f, err = os.Open(input); err != nil {
		return
	}
	defer f.Close()
	var img image.Image
	if img, _, err = image.Decode(f); err != nil {
		return
	}
	fmt.Printf("%s: %v\n", input, img.Bounds())
	return
}

func main() {
	for _, input := range os.Args[1:] {
		if err := HandleImage(input); err != nil {
			log.Fatalf("Unable to handle %s: %v", input)
		}
	}
}
