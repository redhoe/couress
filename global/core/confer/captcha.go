package confer

type Captcha struct {
	KeyLong            int   `mapstructure:"key-long" json:"key-long" yaml:"key-long"`                                     // 验证码长度
	ImgWidth           int   `mapstructure:"img-width" json:"img-width" yaml:"img-width"`                                  // 验证码宽度
	ImgHeight          int   `mapstructure:"img-height" json:"img-height" yaml:"img-height"`                               // 验证码高度
	OpenCaptcha        int   `mapstructure:"open-captcha" json:"open-captcha" yaml:"open-captcha"`                         // 防爆破验证码开启此数，N = 0-10 代表错误N次后出现验证码，<10  不开启
	OpenCaptchaTimeOut int64 `mapstructure:"open-captcha-timeout" json:"open-captcha-timeout" yaml:"open-captcha-timeout"` // 触发防爆破验证码时长，单位：s(分)
}
