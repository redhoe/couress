package modeler

import "golang.org/x/text/language"

type Lang struct {
	Key  language.Tag `json:"lang"`
	Name string       `json:"name"`
}

var Langs = []Lang{
	{language.English, "English"},
	{language.SimplifiedChinese, "中文(简体)"},
	{language.TraditionalChinese, "中文(繁体)"},
	{language.Japanese, "日本語"},
	{language.Korean, "한국어"},
}
