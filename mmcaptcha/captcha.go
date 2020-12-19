package mmcaptcha

import (
	"fmt"
	"image"
	"math/rand"
	"strings"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
)

type mmCaptcha interface {
	GenerateCaptcha() (string, image.Image)
	GenerateLargeCaptcha(c chan *CaptchaPayload)
	SaveCaptchaAsPng(path string) (string, image.Image, string)
	SaveLargeCaptchaAsPng(path string, c chan *CaptchaPayload)
}

type CaptchaPayload struct {
	captchaString string
	image         image.Image
	pathString    string
}

type captcha struct {
	captchautils   CaptchaUtils
	width          float64
	height         float64
	captchaString  []string
	fontFaces      []font.Face
	useNumber      bool
	displayStrokes bool
	strokeCount    int
}

// NewMMCaptcha constructor following parameters
//
// W width : should between 1024 & 512
//
// H height : should between 512 & 256
//
// captchaCount  : number of captcha character you want
//
// useNumber  : captcha will include alphanumeric
//
// displayStrokes  : captcha will display strokes if the value was true
//
// strokeCount  : number of strokes for each captcha character  will display minimum is 0 & maximum is 20
func NewMMCaptcha(W float64, H float64, captchaCount int, useNumber bool, displayStrokes bool, strokeCount int) mmCaptcha {
	utils := NewCaptchaUtils(useNumber)
	capchaArr := utils.GenerateCaptchaString(captchaCount)
	if int(W)%int(H) != 0 || int(W) > 1024 || int(W) < 512 || int(H) > 512 || int(H) < 256 {
		W = 512
		H = 256
	}

	if !displayStrokes {
		displayStrokes = false
	} else {
		displayStrokes = true
	}

	if strokeCount > 20 || strokeCount < 0 {
		strokeCount = 10
	}
	fontSize := H / (float64(len(capchaArr)) / float64(2))
	fontFaceArr := make([]font.Face, len(utils.GetAvailableFonts()))
	for i, v := range utils.GetAvailableFonts() {
		ff, _ := gg.LoadFontFace(fmt.Sprintf("mmcaptcha/%s", v), fontSize)
		fontFaceArr[i] = ff
	}

	return &captcha{
		captchautils:   utils,
		fontFaces:      fontFaceArr,
		captchaString:  capchaArr,
		width:          W,
		height:         H,
		useNumber:      useNumber,
		displayStrokes: displayStrokes,
		strokeCount:    strokeCount,
	}
}

// GenerateCaptcha generate captcha and returns value and image.Image
func (c *captcha) GenerateCaptcha() (string, image.Image) {
	dc := c.drawCaptcha()
	captchaString := strings.Join(c.captchaString, "")
	return captchaString, dc.Image()
}

// GenerateLargeCaptcha generate captcha with goroutine and *CaptchaPayload which include value and image.Image
func (c *captcha) GenerateLargeCaptcha(capchaChan chan *CaptchaPayload) {
	go func() {
		dc := c.drawCaptcha()
		captchaString := strings.Join(c.captchaString, "")
		capchaChan <- &CaptchaPayload{
			captchaString: captchaString,
			image:         dc.Image(),
		}

	}()
}

// SaveCaptchaAsPng generate captcha and save to provided path and return  value ,image.Image, outputpath
func (c *captcha) SaveCaptchaAsPng(path string) (string, image.Image, string) {

	dc := c.drawCaptcha()
	captchaString := strings.Join(c.captchaString, "")

	if path == "" {
		path = captchaString
	}
	pathString := fmt.Sprintf("%s.png", path)
	dc.SavePNG(pathString)
	return captchaString, dc.Image(), pathString
}

// SaveLargeCaptchaAsPng generate captcha and save to provided path with goroutine and return *CaptchaPayload with value ,image.Image, outputpath
func (c *captcha) SaveLargeCaptchaAsPng(path string, capchaChan chan *CaptchaPayload) {

	go func() {
		dc := c.drawCaptcha()
		captchaString := strings.Join(c.captchaString, "")
		if path == "" {
			path = captchaString
		}
		pathString := fmt.Sprintf("%s.png", path)
		dc.SavePNG(pathString)
		capchaChan <- &CaptchaPayload{
			captchaString: captchaString,
			image:         dc.Image(),
			pathString:    pathString,
		}

	}()

}

// drawCaptcha draws captcha image
func (c *captcha) drawCaptcha() *gg.Context {

	var W = int(c.width)
	var H = int(c.height)
	dc := gg.NewContext(int(c.width), int(c.height))

	dc.SetRGB(c.captchautils.GenerateRandomColor())
	dc.DrawRectangle(0, 0, float64(W), float64(H))
	dc.Fill()

	for i, v := range c.captchaString {
		fontSize := float64(H) / (float64(len(c.captchaString)) / float64(2))
		dc.SetRGB(c.captchautils.GenerateRandomColor())

		if (i+1)%2 == 0 {
			dc.SetFontFace(c.fontFaces[c.captchautils.GenerateRandomNumber(len(c.fontFaces)-1, 0)])
			dc.DrawStringAnchored(v, fontSize*float64(i+1)-float64(H/len(c.captchaString)), float64(H)/2, 0.5, 0.5)
		} else {
			dc.SetFontFace(c.fontFaces[c.captchautils.GenerateRandomNumber(len(c.fontFaces)-1, 0)])
			dc.DrawStringAnchored(v, fontSize*float64(i+1)-float64(H/len(c.captchaString)), float64(H)/2, 0.5, 0.5)
		}
		if c.displayStrokes {
			c.drawLines(dc, float64(W), float64(H), c.strokeCount)
		}

	}
	return dc
}

// Draw Lines for captcha image visibility check
func (c *captcha) drawLines(dc *gg.Context, W float64, H float64, strokeCount int) {
	for i := 0; i < strokeCount; i++ {
		x1 := rand.Float64() * W
		y1 := rand.Float64() * H
		x2 := rand.Float64() * W
		y2 := rand.Float64() * H
		r := rand.Float64()
		g := rand.Float64()
		b := rand.Float64()
		a := rand.Float64()*0.5 + 0.5
		w := rand.Float64()*4 + 1
		dc.SetRGBA(r, g, b, a)
		dc.SetLineWidth(w)
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
	}
}
