package tips

// 使用指定变量名来进行多语言

func init() {
	// 状态码
	cn["s_success"] = "成功"
	cn["s_fail"] = "失败"

	// 变量
	cn["v_title"] = `3天行权资产包`
	en["v_title"] = `3-day exercise asset package`
	kr["v_title"] = "3일 행정권 자산 패키지"
}

// SetVariableMultilingual 用于多语言字典写入缓存
func SetVariableMultilingual(lang, key, value string) {
	langMap := getMapByLang(lang)
	langMap[key] = value
}
