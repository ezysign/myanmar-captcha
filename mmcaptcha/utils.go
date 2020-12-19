package mmcaptcha

import (
	"math/rand"
	"time"
)

type CaptchaUtils interface {
	GenerateRandomNumber(max int, min int) int
	GenerateRandomDecimalNumber(max int, min int) float64
	GenerateRandomCharacter() string
	GenerateRandomColor() (float64, float64, float64)
	GenerateAlphaNumericCharacters() string
	GenerateCaptchaString(count int) []string
	GetRandomFont() string
	GetAvailableFonts() []string
}

type captchaUtils struct {
	useNumber  bool
	characters []string
	numbers    []string
	Fonts      []string
}

func NewCaptchaUtils(useNumber bool) CaptchaUtils {
	return &captchaUtils{
		useNumber:  useNumber,
		characters: []string{"က", "ခ", "ဂ", "ဃ", "င", "စ", "ဆ", "ဇ", "ဈ", "ည", "ဋ", "ဌ", "ဍ", "ဎ", "ဏ", "တ", "ထ", "ဒ", "ဓ", "န", "ပ", "ဖ", "ဗ", "ဘ", "မ", "ယ", "ရ", "လ", "ဝ", "သ", "ဟ", "အ"},
		numbers:    []string{"၁", "၂", "၃", "၄", "၅", "၆", "၇", "၈", "၉", "၀"},
		Fonts:      []string{"Padauk-Bold.ttf", "Padauk-Regular.ttf"},
	}
}

func (c *captchaUtils) GenerateRandomNumber(max int, min int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func (c *captchaUtils) GenerateRandomDecimalNumber(max int, min int) float64 {
	rand.Seed(time.Now().UnixNano())
	return float64(rand.Intn(max-min+1) + min)
}

func (c *captchaUtils) GenerateRandomCharacter() string {

	return c.characters[c.GenerateRandomNumber(len(c.characters)-1, 0)]
}

func (c *captchaUtils) GenerateAlphaNumericCharacters() string {
	arr := append(c.characters, c.numbers...)
	return arr[c.GenerateRandomNumber(len(arr)-1, 0)]
}

func (c *captchaUtils) GenerateRandomColor() (float64, float64, float64) {
	return c.GenerateRandomDecimalNumber(255, 0), c.GenerateRandomDecimalNumber(255, 0), c.GenerateRandomDecimalNumber(255, 0)

}

func (c *captchaUtils) GenerateCaptchaString(count int) []string {
	arr := make([]string, count)
	for i := range arr {
		if c.useNumber {
			arr[i] = c.GenerateAlphaNumericCharacters()
		} else {
			arr[i] = c.GenerateRandomCharacter()
		}

	}
	return arr
}

func (c *captchaUtils) GetRandomFont() string {
	return c.Fonts[c.GenerateRandomNumber(len(c.Fonts)-1, 0)]
}
func (c *captchaUtils) GetAvailableFonts() []string {
	return c.Fonts
}
