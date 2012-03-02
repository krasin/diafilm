package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"log"
	"math"
	"os"
)

const maxScore = 0.004
const minColor = 0

func sqr(x float64) float64 {
	return x * x
}

func sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func ns(x uint32) float64 {
	if x < minColor {
		x = 0
	}
	return sqr(float64(x) / 65536)
}

func max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func l2(c color.Color) float64 {
	r, g, b, _ := c.RGBA()
	return max(ns(r), max(ns(g), ns(b)))
}

func scoreYMin(img image.Image, min, max image.Point) (score float64) {
	for x := min.X; x < max.X; x++ {
		score += l2(img.At(x, min.Y))
	}
	score = sqrt(score)
	score /= float64(max.X - min.X)
	fmt.Printf("scoreYMin(%v, %v): %v\n", min, max, score)
	return
}

func scoreYMax(img image.Image, min, max image.Point) (score float64) {
	for x := min.X; x < max.X; x++ {
		score += l2(img.At(x, max.Y-1))
	}
	score = sqrt(score)
	score /= float64(max.X - min.X)
	fmt.Printf("scoreYMax(%v, %v): %v\n", min, max, score)
	return
}

func scoreXMin(img image.Image, min, max image.Point) (score float64) {
	for y := min.Y; y < max.Y; y++ {
		score += l2(img.At(min.X, y))
	}
	score = sqrt(score)
	score /= float64(max.Y - min.Y)
	fmt.Printf("scoreXMin(%v, %v): %v\n", min, max, score)
	return
}

func scoreXMax(img image.Image, min, max image.Point) (score float64) {
	for y := min.Y; y < max.Y; y++ {
		score += l2(img.At(max.X-1, y))
	}
	score = sqrt(score)
	score /= float64(max.Y - min.Y)
	fmt.Printf("scoreXMax(%v, %v): %v\n", min, max, score)
	return
}

func findActualImage(img image.Image) (r image.Rectangle) {
	fmt.Printf("findActualImage\n")
	min := img.Bounds().Min
	max := img.Bounds().Max

	cur := scoreXMin(img, min, max)
	for cur < maxScore && min.X < max.X {
		min.X++
		cur = scoreXMin(img, min, max)
	}

	cur = scoreXMax(img, min, max)
	for cur < maxScore && min.X < max.X {
		max.X--
		cur = scoreXMax(img, min, max)
	}

	cur = scoreYMin(img, min, max)
	for cur < maxScore && min.Y < max.Y {
		min.Y++
		cur = scoreYMin(img, min, max)
	}

	cur = scoreYMax(img, min, max)
	for cur < maxScore && min.Y < max.Y {
		max.Y--
		cur = scoreYMax(img, min, max)
	}

	return image.Rectangle{min, max}
}

func handleImage(input string) (err error) {
	var f *os.File
	if f, err = os.Open(input); err != nil {
		return
	}
	defer f.Close()
	var img image.Image
	if img, _, err = image.Decode(f); err != nil {
		return
	}
	r := findActualImage(img)
	fmt.Printf("%s: %v, actual image is inside %v\n", input, img.Bounds(), r)
	return
}

func main() {
	for _, input := range os.Args[1:] {
		if err := handleImage(input); err != nil {
			log.Fatalf("Unable to handle %s: %v", input)
		}
	}
}
