package main

import (
	"image"
	"image/draw"

	"github.com/echo-music/go-image/marker"

	"github.com/disintegration/imaging"
)

func main() {
	var imgObj = marker.NewImage()
	bgImg, err := imgObj.Background("../images/bg.png")
	if err != nil {
		panic(err)
	}

	// 头像设置
	petRoleImg, err := imgObj.Circle("../images/dog.png", 180)
	if err != nil {
		panic(err)
	}
	userImg, err := imgObj.Circle("../images/girl.png", 80)
	if err != nil {
		panic(err)
	}

	// 坐标设置
	draw.Draw(bgImg, petRoleImg.Bounds().Add(image.Pt(18, 15)), petRoleImg, image.Point{}, draw.Over)
	draw.Draw(bgImg, userImg.Bounds().Add(image.Pt(3, 123)), userImg, image.Point{}, draw.Over)

	// 保存到临时文件
	filepath := "../images/marker.png"
	err = imaging.Save(bgImg, filepath)
	if err != nil {
		panic(err)
	}

}
