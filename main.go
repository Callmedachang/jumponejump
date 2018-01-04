package main

import (
	"image/jpeg"
	"log"
	"os"
	"fmt"
)

func main() {
	// open ".jpg"
	file, err := os.Open("src_file/IMG_3326.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	m, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	// and preserve aspect ratio
	maxX := m.Bounds().Max.X
	maxY := m.Bounds().Max.Y
LOOP:
	for j := 600; j < maxY; j++ {
		//背景色RGB
		zeroPointR, zeroPointG, zeroPointB, _ := m.At(maxX-3, j).RGBA()
		start := 0
		isBackground := true
		for i := maxX - 1; i > 1; i-- {
			r, g, b, _ := m.At(i, j).RGBA()
			if r == zeroPointR && g == zeroPointG && b == zeroPointB {
				continue
			} else {
				if isBackground {
					start = i
					isBackground = false
				}
				if i <= start {
					fmt.Println("中心点:(", start, ",", j, ")")
					break LOOP
				}
			}
		}
	}
}
