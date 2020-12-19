package main

import (
	"fmt"
	"myanmar-captcha/mmcaptcha"
)

func main() {
	captcha := mmcaptcha.NewMMCaptcha(1024, 512, 5, false, true, 15)
	cap := make(chan *mmcaptcha.CaptchaPayload)
	captcha.SaveLargeCaptchaAsPng("out", cap)
	payload := <-cap
	fmt.Println(payload)
	close(cap)

}
