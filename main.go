package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"github.com/nfnt/resize"
	"math"
	"jumponejump/source"
)

func main() {
	guyPositionX, guyPositionY, targetPositionX, targetPositionY,len:=source.DealImage("src_file/IMG_3326.jpg")
	fmt.Println("目标点坐标:(",targetPositionX,",",targetPositionY,")")
	fmt.Println("人物中心坐标:(",guyPositionX,",",guyPositionY,")")
	fmt.Println("跳跃距离(像素):",len)
}



