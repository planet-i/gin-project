package article_service

import (
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"os"

	"github.com/golang/freetype"

	"github.com/planet-i/gin-project/pkg/file"
	"github.com/planet-i/gin-project/pkg/qrcode"
	"github.com/planet-i/gin-project/pkg/setting"
)

type ArticlePoster struct {
	PosterName string
	*Article
	Qr *qrcode.QrCode
}

func NewArticlePoster(posterName string, article *Article, qr *qrcode.QrCode) *ArticlePoster {
	return &ArticlePoster{
		PosterName: posterName,
		Article:    article,
		Qr:         qr,
	}
}

func GetPosterFlag() string {
	return "poster"
}

func (a *ArticlePoster) CheckMergedImage(path string) bool {
	if file.CheckNotExist(path+a.PosterName) == true {
		return false
	}

	return true
}

func (a *ArticlePoster) OpenMergedImage(path string) (*os.File, error) {
	f, err := file.MustOpen(a.PosterName, path)
	if err != nil {
		return nil, err
	}

	return f, nil
}

type ArticlePosterBg struct {
	Name string
	*ArticlePoster
	*Rect
	*Pt
}

type Rect struct {
	Name string
	X0   int
	Y0   int
	X1   int
	Y1   int
}

type Pt struct {
	X int
	Y int
}

func NewArticlePosterBg(name string, ap *ArticlePoster, rect *Rect, pt *Pt) *ArticlePosterBg {
	return &ArticlePosterBg{
		Name:          name,
		ArticlePoster: ap,
		Rect:          rect,
		Pt:            pt,
	}
}

type DrawText struct {
	JPG    draw.Image
	Merged *os.File

	Title string
	X0    int
	Y0    int
	Size0 float64

	SubTitle string
	X1       int
	Y1       int
	Size1    float64
}

func (a *ArticlePosterBg) DrawPoster(d *DrawText, fontName string) error {
	fontSource := setting.AppSetting.RuntimeRootPath + setting.AppSetting.FontSavePath + fontName
	fontSourceBytes, err := ioutil.ReadFile(fontSource)
	if err != nil {
		return err
	}

	trueTypeFont, err := freetype.ParseFont(fontSourceBytes)
	if err != nil {
		return err
	}

	fc := freetype.NewContext()
	fc.SetDPI(72)
	fc.SetFont(trueTypeFont)
	fc.SetFontSize(d.Size0)
	fc.SetClip(d.JPG.Bounds())
	fc.SetDst(d.JPG)
	fc.SetSrc(image.Black)

	pt := freetype.Pt(d.X0, d.Y0)
	_, err = fc.DrawString(d.Title, pt)
	if err != nil {
		return err
	}

	fc.SetFontSize(d.Size1)
	_, err = fc.DrawString(d.SubTitle, freetype.Pt(d.X1, d.Y1))
	if err != nil {
		return err
	}

	err = jpeg.Encode(d.Merged, d.JPG, nil)
	if err != nil {
		return err
	}

	return nil
}

func (a *ArticlePosterBg) Generate() (string, string, error) {
	// 获取二维码存储路径
	fullPath := qrcode.GetQrCodeFullPath()
	fileName, path, err := a.Qr.Encode(fullPath)
	if err != nil {
		return "", "", err
	}
	// 检查合并后图像（指的是存放合并后的海报）是否存在
	if !a.CheckMergedImage(path) {
		// 生成待合并的图像 mergedF
		mergedF, err := a.OpenMergedImage(path)
		if err != nil {
			return "", "", err
		}
		defer mergedF.Close()
		// 打开事先存放的背景图 bgF
		bgF, err := file.MustOpen(a.Name, path)
		if err != nil {
			return "", "", err
		}
		defer bgF.Close()
		// 打开生成的二维码图像 qrF
		qrF, err := file.MustOpen(fileName, path)
		if err != nil {
			return "", "", err
		}
		defer qrF.Close()
		// 解码 bgF 和 qrF 返回 image.Image
		bgImage, err := jpeg.Decode(bgF)
		if err != nil {
			return "", "", err
		}
		// 创建一个新的 RGBA 图像
		qrImage, err := jpeg.Decode(qrF)
		if err != nil {
			return "", "", err
		}
		// 在 RGBA 图像上绘制 背景图（bgF）
		jpg := image.NewRGBA(image.Rect(a.Rect.X0, a.Rect.Y0, a.Rect.X1, a.Rect.Y1))
		// 在已绘制背景图的 RGBA 图像上，在指定 Point 上绘制二维码图像（qrF）
		draw.Draw(jpg, jpg.Bounds(), bgImage, bgImage.Bounds().Min, draw.Over)
		draw.Draw(jpg, jpg.Bounds(), qrImage, qrImage.Bounds().Min.Sub(image.Pt(a.Pt.X, a.Pt.Y)), draw.Over)
		// 在图片上绘制文字
		err = a.DrawPoster(&DrawText{
			JPG:    jpg,
			Merged: mergedF,

			Title: "Golang Gin 系列文章",
			X0:    80,
			Y0:    160,
			Size0: 42,

			SubTitle: "---煎鱼",
			X1:       320,
			Y1:       220,
			Size1:    36,
		}, "msyhbd.ttc")

		if err != nil {
			return "", "", err
		}
	}

	return fileName, path, nil
}
