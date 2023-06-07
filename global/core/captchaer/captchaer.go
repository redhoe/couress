package captchaer

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/redhoe/couress/utils/simple"
	"image/color"
	"image/png"
	"strings"
)

type CaptchaEngine struct {
}

var cap *captcha.Captcha
var capStore *hashmap.Map

func NewCaptchaEngine() *CaptchaEngine {
	if cap == nil {
		cap = captcha.New()
		// 设置字体
		_ = cap.SetFont("./comic.ttf")
		// 设置验证码大小
		cap.SetSize(128, 64)
		// 设置干扰强度
		cap.SetDisturbance(captcha.MEDIUM)
		// 设置前景色 可以多个 随机替换文字颜色 默认黑色
		cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
		// 设置背景色 可以多个 随机替换背景色 默认白色
		cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	}
	if capStore == nil {
		capStore = hashmap.New()
	}
	return &CaptchaEngine{}
}

func (c *CaptchaEngine) Generate() (bs64 string, cid string) {
	img, code := c.getImg()
	cid = base64.StdEncoding.EncodeToString([]byte(simple.NewUuid()))
	// todo: set cid to store
	capStore.Put(cid, fmt.Sprintf("%s", code))
	w := bytes.NewBuffer(nil)
	_ = png.Encode(w, img)
	return fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(w.Bytes())), cid
}

func (c *CaptchaEngine) Verify(cid, answer string) (match bool) {
	vv, ok := capStore.Get(cid)
	if !ok {
		return ok
	}
	v, _ := vv.(string)
	vv = strings.TrimSpace(v)
	capStore.Remove(cid)
	return vv == strings.TrimSpace(answer)
}

func (*CaptchaEngine) getImg() (*captcha.Image, string) {
	// 创建验证码 4个字符 captcha.NUM 字符模式数字类型
	// 返回验证码图像对象以及验证码字符串 后期可以对字符串进行对比 判断验证
	return cap.Create(4, captcha.CLEAR)
}
