package simple

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func Str2sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func Struct2MapByJson(info any) map[string]any {
	data, _ := json.Marshal(&info)
	m := make(map[string]any)
	_ = json.Unmarshal(data, &m)
	return m
}

func MapToStruct[T any](req map[string]any, t T) (T, error) {
	data, _ := json.Marshal(req)
	err := json.Unmarshal(data, &t)
	return t, err
}

// Convert map json string

func MapToJson(m map[string]string) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Marshal with error: %+v\n", err)
		return "", nil
	}
	return string(jsonByte), nil
}

// Convert json string to map

func JsonToMap(jsonStr string) (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}

	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}

	return m, nil
}

func Struct2StringPrt(c any) *string {
	b, _ := json.Marshal(c)
	bt := string(b)
	return &bt
}

func Struct2String(c any) string {
	b, _ := json.Marshal(c)
	return string(b)
}

// SnakeString 驼峰转蛇形 XxYy to xx_yy , XxYY to xx_y_y
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

func SnakeMap(m map[string]any) map[string]any {
	mt := make(map[string]any)
	for k, v := range m {
		mt[SnakeString(k)] = v
	}
	return mt
}
