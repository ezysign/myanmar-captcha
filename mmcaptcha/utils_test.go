package mmcaptcha

import (
	"sort"
	"testing"
)

var characters = []string{"က", "ခ", "ဂ", "ဃ", "င", "စ", "ဆ", "ဇ", "ဈ", "ည", "ဋ", "ဌ", "ဍ", "ဎ", "ဏ", "တ", "ထ", "ဒ", "ဓ", "န", "ပ", "ဖ", "ဗ", "ဘ", "မ", "ယ", "ရ", "လ", "ဝ", "သ", "ဟ", "အ"}
var numbers = []string{"၁", "၂", "၃", "၄", "၅", "၆", "၇", "၈", "၉", "၀"}
var fonts = []string{"Padauk-Bold.ttf", "Padauk-Regular.ttf"}

func TestRandomCharacter(t *testing.T) {
	captchaUtils := NewCaptchaUtils(false)

	if sort.SearchStrings(characters, captchaUtils.GenerateRandomCharacter()) == -1 {
		t.Log("it Should be random Character")
		t.Fail()
	}

}

func TestRandomAlphaNumericCharacter(t *testing.T) {
	captchaUtils := NewCaptchaUtils(true)
	alphanumericArr := append(characters, numbers...)
	if sort.SearchStrings(alphanumericArr, captchaUtils.GenerateAlphaNumericCharacters()) == -1 {
		t.Log("it Should be random alphanumeric Character")
		t.Fail()
	}
}

func TestGenerateCaptchaString(t *testing.T) {
	captchaUtils := NewCaptchaUtils(false)

	for _, v := range captchaUtils.GenerateCaptchaString(3) {
		if sort.SearchStrings(characters, v) == -1 {
			t.Log("it Should be random  Character")
			t.Fail()
		}

	}
}

func TestGenerateAlphaNumericCaptchaString(t *testing.T) {
	captchaUtils := NewCaptchaUtils(true)
	alphanumericArr := append(characters, numbers...)
	for _, v := range captchaUtils.GenerateCaptchaString(4) {
		if sort.SearchStrings(alphanumericArr, v) == -1 {
			t.Log("it Should be random  AlphaNumeric Character")
			t.Fail()
		}

	}
}

func TestRandomFont(t *testing.T) {
	captchaUtils := NewCaptchaUtils(true)
	if sort.SearchStrings(fonts, captchaUtils.GetRandomFont()) == -1 {
		t.Log("it Should be random  Font face")
		t.Fail()
	}

}
