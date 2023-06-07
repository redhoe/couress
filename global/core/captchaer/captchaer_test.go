package captchaer

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"testing"
)

func TestGetCaptcha(t *testing.T) {
	img, str := NewCaptchaEngine().getImg()
	w := bytes.NewBuffer(nil)
	_ = png.Encode(w, img)
	//t.Log(w.String())
	t.Log("data:image/png;base64," + base64.StdEncoding.EncodeToString(w.Bytes()))
	t.Log(str)
}
