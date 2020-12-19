package main

import (
	"fmt"

	"github.com/ezysign/myanmar-captcha/mmcaptcha"
)

func main() {
	captcha := mmcaptcha.NewMMCaptcha(1024, 512, 5, false, true, 15)
	cap := make(chan *mmcaptcha.CaptchaPayload)
	captcha.GenerateLargeCaptcha(cap)
	payload := <-cap
	fmt.Println(payload)
	close(cap)

}
