package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	//图片大小
	const size = 100
	//创建灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))

	//将底色填充为白色
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{Y: 255})
		}
	}
	//计算每个sin点在什么地方,并将那个地方的点填为黑色
	for x := 0; x < size; x++ {
		s := float64(x) * 2 * math.Pi / size
		y := size/2 - math.Sin(s)*size/2
		pic.SetGray(x, int(y), color.Gray{Y: 0})
	}

	//创建一个文件
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}

	//将png格式的数据写入文件中
	_ = png.Encode(file, pic)

	//关闭文件
	_ = file.Close()
}
