package simple

import (
	"fmt"
	"regexp"
)

// 验证器

// 密码强度：（长度大于等于N位,包含⼤⼩写字⺟+数字）

func CheckPasswordLever(ps string) error {
	if len(ps) < 6 {
		return fmt.Errorf("len is < 6")
	}
	return nil
}

// 验证邮箱格式

func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 验证手机格式

func VerifyMobileFormat(mobile string) bool {
	pattern := `^[1-9][0-9]{5,15}$` // 匹配手机号码
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(mobile)
}

// 密码包含各类字符检查

func CheckPasswordStrong(ps string) error {
	if len(ps) < 6 {
		return fmt.Errorf("password len is < 6")
	}
	num := `[0-9]{1}`
	az := `[a-z]{1}`
	AZ := `[A-Z]{1}`
	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		return fmt.Errorf("password need num :%v", err)
	}
	if b, err := regexp.MatchString(az, ps); !b || err != nil {
		return fmt.Errorf("password need a_z :%v", err)
	}
	if b, err := regexp.MatchString(AZ, ps); !b || err != nil {
		return fmt.Errorf("password need A_Z :%v", err)
	}
	symbol := `[!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(symbol, ps); !b || err != nil {
		return fmt.Errorf("password need symbol :%v", err)
	} // 特殊字符
	return nil
}
