package marker

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/disintegration/imaging"
)

// 圆形遮罩
type circle struct {
	p image.Point // 圆心位置
	r int         // 半径
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{A: 255} // 半径以内的图案设成完全不透明
	}
	return color.Alpha{}
}

type Image struct {
}

func NewImage() *Image {
	return &Image{}
}

// 生成图片背景
func (img *Image) Background(filepath string) (*image.RGBA, error) {

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
}

// 生成圆状图片
func (img *Image) Circle(filepath string, l int) (image.Image, error) {

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
}
