
使用go将多张图片合并成一张图片，能做到么？先说答案，当然可以咯。最近兼职了一个项目，客户要求给你三张图(背景图,用户头像,宠物头像)合并成一张图（地图标记）
比如要合并成下面地图图标,如何去实现呢？

[![pkJz1Rs.png](https://s21.ax1x.com/2024/06/05/pkJz1Rs.png)](https://imgse.com/i/pkJz1Rs)


# 快速体验
```
git clone git@github.com:echo-music/go-image.git

cd go-image && go mod download

cd marker/cmd && go run main.go
```

# 快速入门

## 1、绘制圆形图片
marker.go
代码如下：
```
    // 打开源图片
    imgSource, err := imaging.Open(filepath)
	if err != nil {
		return nil, err
	}
	if imgSource == nil {
		return nil, errors.New("打开图片资源失败")
	}
	width := imgSource.Bounds().Dx()
	height := imgSource.Bounds().Dy()
	size := width
	if height < width {
		size = height
	}

	avatarRad := size / 2
	c := circle{p: image.Point{X: avatarRad, Y: avatarRad}, r: avatarRad}
	circleAvatar := image.NewRGBA(image.Rect(0, 0, avatarRad*2, avatarRad*2))
	// DrawMask 函数可以在 src 上面一个遮罩，可以实现圆形图片、圆角等效果。
	draw.DrawMask(circleAvatar, circleAvatar.Bounds(), imgSource, image.Point{X: (width-size)/2, Y: (height-size)/2}, &c, image.Point{}, draw.Over) // 使用 Over 模式进行混合

	//返回缩放比例后的圆形图片
	return imaging.Resize(circleAvatar, l, l, imaging.Lanczos), nil
```


## 2、绘制背景图片
marker.go
```
    var (
		imgSource image.Image
		err       error
	)

	imgSource, err = imaging.Open(filepath)
	if err != nil {
		return nil, err
	}
	if imgSource == nil {
		return nil, fmt.Errorf("Background:打开图片资源失败")
	}

	// 创建一个新的RGBA格式的图片，大小与背景图片相同
	bounds := imgSource.Bounds()
	bgImg := image.NewRGBA(bounds)

	// 将背景图片绘制到新的RGBA
	draw.Draw(bgImg, bounds, imgSource, image.Point{0, 0}, draw.Src)
	return bgImg, nil
```


## 3、合并图片
cmd/main.go
```
package main

import (
	"image"
	"image/draw"

	"github.com/disintegration/imaging"
	"github.com/echo-music/go-learn/marker"
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

```
