package filer

import (
	"fmt"
	"github.com/redhoe/couress/global/core/confer"
	"github.com/redhoe/couress/utils/simple"
	"github.com/samber/lo"
	"mime/multipart"
	"strings"
	"time"
)

// UploadMultipart 文件流上传 根据配置上传文件 获得url
func UploadMultipart(file *multipart.FileHeader, args ...string) (string, error) {
	// 尺寸检查
	if err := sizeCheck(file); err != nil {
		return "", err
	}
	// 类型检查
	fileName, err := typeCheck(file)
	if err != nil {
		return "", err
	}

	// 保存路径
	argsNew := make([]string, 0)
	argsNew = append(argsNew, localStorePath())
	argsNew = append(argsNew, time.Now().Format("2006-01-02"))
	argsNew = append(argsNew, args...)

	// 保存本地访问路径:加入传参
	localStorePathExtend := simple.AbsPath(argsNew...)
	visitPath := simple.ToPath(localUrl(), simple.ToPath(argsNew...))

	// 获得配置文件 oss/local 默认本地
	ossType := strings.ToLower(confer.AppConfServer.System.OssType)
	switch ossType {
	case "ali":
		return AliEngin().UploadMultipart(file, fileName, simple.ToPath(argsNew...))
	case "local":
		return LocalEngin().UploadMultipart(file, fileName, localStorePathExtend, visitPath)
	default:
		return LocalEngin().UploadMultipart(file, fileName, localStorePathExtend, visitPath)
	}
}

func localUrl() string {
	return confer.AppConfServer.Local.Path
}

func localStorePath() string {
	return confer.AppConfServer.Local.StorePath
}

func sizeCheck(file *multipart.FileHeader) error {
	maxSize := confer.AppConfServer.Local.MaxSize
	if file.Size >= 1024*1024*maxSize {
		return fmt.Errorf("max file size:%dM", maxSize)
	}
	return nil
}

func typeCheck(file *multipart.FileHeader) (string, error) {
	a := strings.Split(file.Filename, ".")
	fileType := strings.ToLower(a[len(a)-1])
	types := strings.Split(confer.AppConfServer.Local.Types, ",")
	if !lo.Contains(types, fileType) {
		return "", fmt.Errorf("SupportFileType:%s", fileType)
	}

	return simple.ToPath(simple.NewFileName(fileType)), nil
}
