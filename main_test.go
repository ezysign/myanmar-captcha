package main

import (
	"testing"

	"github.com/ezysign/myanmar-captcha/mmcaptcha"
)

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		captcha := mmcaptcha.NewMMCaptcha(1024, 512, 5, true, true, 15)
		cap := make(chan *mmcaptcha.CaptchaPayload)
		captcha.SaveLargeCaptchaAsPng("out", cap)
		<-cap

	}
}
