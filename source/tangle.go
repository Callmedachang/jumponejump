package source

import (
	"os"
	"image/jpeg"
	"github.com/nfnt/resize"
	"math"
	"image"
	"log"
)

func DealImage(path string) (guyPositionX, guyPositionY, targetPositionX, targetPositionY float32,len float64) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	// decode jpeg into image.Image
	m, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	times := float32(m.Bounds().Max.X) / 1000.00
	//缩放
	m = resize.Resize(1000, 0, m, resize.Lanczos3)
	//任务位置
	guyPositionX, guyPositionY = getGuyPosition(m)
	if int(guyPositionX) > (m.Bounds().Max.X / 2) {
		//人物在右边
		targetPositionX, targetPositionY = getCenterPointRight(m)
	} else { //人物在左边
		targetPositionX, targetPositionY = getCenterPointLeft(m)
	}
	temp:=math.Abs(float64(targetPositionX-guyPositionX)*float64(targetPositionX-guyPositionX))+
		math.Abs(float64(targetPositionY-guyPositionY)*float64(targetPositionY-guyPositionY))
	len=math.Sqrt(temp)
	guyPositionX, guyPositionY, targetPositionX, targetPositionY=guyPositionX*times, guyPositionY*times, targetPositionX*times, targetPositionY*times
	len=len*float64(times)
	file.Close()
	return
}

func getCenterPointLeft(m image.Image) (x, y float32) {
	maxX := m.Bounds().Max.X
	maxY := m.Bounds().Max.Y
	len := 0
LOOP:
	for j := 600; j < maxY; j = j + 2 {
		//背景色RGB
		zeroPointR, zeroPointG, zeroPointB, _ := m.At(maxX-3, j).RGBA()
		start, end := 0, 0
		isBackground := true
		for i := 0; i < maxX-1; i++ {
			//当前像素点RGB
			r, g, b, _ := m.At(i, j).RGBA()
			if r == zeroPointR && g == zeroPointG && b == zeroPointB {
				if !isBackground {
					end = i
					if end > len {
						len = end
						isBackground = true
						break
					} else {
						x = float32((end-start)/2 + start)
						y = float32(j)
						if x > 200 {
							break LOOP
						}
					}
				}
			} else {
				if isBackground {
					start = i
					isBackground = false
				}
			}
		}
	}
	return
}

func getCenterPointRight(m image.Image) (x, y float32) {
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
					x = float32(start)
					y = float32(j)
					break LOOP
				}
			}
		}
	}
	return
}

func getGuyPosition(m image.Image) (x, y float32) {
	maxX := m.Bounds().Max.X
	maxY := m.Bounds().Max.Y
	lenOut, startOut, outJ := 0, 0, 0
	error := float64(1560)
	for j := 600; j < maxY; j = j + 2 {
		//跳动人的RGB
		var zeroPointR, zeroPointG, zeroPointB = uint32(54*256), uint32(52*256), uint32(92*256)
		end, start := 0, 0
		isBackground := true
		for i := 0; i < maxX-1; i++ {
			//当前像素点RGB
			r, g, b, _ := m.At(i, j).RGBA()
			//跳动小人的RGB判断范围误差10个色度
			if abs(r-zeroPointR) < error && abs(g-zeroPointG) < error && abs(b-zeroPointB) < error {
				if !isBackground {
					end = i
					if end-start > lenOut {
						lenOut = end - start
						startOut = start
						outJ = j
						isBackground = true
						break
					}
				}
			} else {
				if isBackground {
					start = i
					isBackground = false
				}
			}
		}
	}
	x = float32(startOut + lenOut/2)
	y = float32(outJ)
	return
}

func abs(number uint32) float64 {
	return math.Abs(float64(number))
}