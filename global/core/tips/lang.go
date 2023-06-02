package tips

import (
	"fmt"
	"github.com/samber/lo"
	"strings"
)

type LangType string

var cn = map[string]string{}
var en = map[string]string{}
var kr = map[string]string{}

const (
	LANGChinese LangType = "zh-cn"
	LANGEnglish LangType = "en-us"
	LANGKorean  LangType = "rp-kr"
)

func LangList() []LangType {
	return []LangType{LANGChinese, LANGEnglish, LANGKorean}
}

func (c LangType) check() bool {
	chainArray := LangList()
	_, ok := lo.Find(chainArray, func(item LangType) bool {
		return item == c
	})
	return ok
}

func (l LangType) String() string {
	return string(l)
}

var DefaultLang = LANGChinese

// 系统公共提示信息

// 定义语种 en-US、zh-CN、rp-KR
var langMap = map[LangType]map[string]string{
	LANGChinese: cn,
	LANGEnglish: en,
	LANGKorean:  kr,
}

func getMapByLang(lang string) map[string]string {
	langType := strToLang(lang)
	switch langType {
	case LANGChinese:
		return cn
	case LANGEnglish:
		return en
	case LANGKorean:
		return kr
	default:
		return en
	}
}

func strToLang(lang string) LangType {
	lang = strings.ToLower(lang)
	return lo.If(LangType(lang).check(), LangType(lang)).Else(DefaultLang)
}

func getLang(value string, langKey string, args ...any) string {
	if value[:2] == "f_" {
		if msg, ok := langMap[strToLang(langKey)][value]; ok {
			return fmt.Sprintf(msg, args)
		}
	}
	if msg, ok := langMap[strToLang(langKey)][value]; ok {
		return msg
	}
	return value
}

func GetLang(value string, langKey string, args ...any) string {
	if value[:2] == "f_" {
		if msg, ok := langMap[strToLang(langKey)][value]; ok {
			return fmt.Sprintf(msg, args)
		}
	}
	if msg, ok := langMap[strToLang(langKey)][value]; ok {
		return msg
	}
	return value
}
