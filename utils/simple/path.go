package simple

import (
	"os"
	"strings"
)

// AbsPath 绝对路径生成
func AbsPath(args ...string) string {
	absPath, err := os.Getwd()
	if err != nil {
		return ""
	}
	savePath := ToPath(absPath, ToPath(args...))
	err = CheckPathIsExist(savePath)
	if err != nil {
		return ""
	}
	return savePath
}

// ToPath 转化成路径
func ToPath(args ...string) string {
	if len(args) == 0 {
		return ""
	}
	strList := make([]string, 0)
	for i, p := range args {
		if strings.TrimSpace(p) == "" {
			continue
		}
		if i == 0 {
			strList = append(strList, strings.TrimRight(strings.TrimSpace(p), "/"))
			continue
		}

		strList = append(strList, strings.Trim(strings.TrimSpace(p), "/"))
	}
	return strings.Join(strList, "/")
}

// CheckPathIsExist 检查路径并创建
func CheckPathIsExist(filePath string) error {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(filePath, os.ModePerm) //生成多级目录
			if err != nil {
				return err
			}
			return err
		}
	}
	return err
}
