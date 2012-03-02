package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
)

func main() {
	for _, input := range os.Args[1:] {
		f, err := os.Open(input)
		if err != nil {
			log.Fatalf("Unable to open input file %s: %v", input, err)
		}
		var img image.Image
		if img, _, err = image.Decode(f); err != nil {
			log.Fatalf("Unable to parse image from %s: %v", input, err)
		}
		f.Close()
		fmt.Printf("%s: %v\n", input, img.Bounds())
	}
}
