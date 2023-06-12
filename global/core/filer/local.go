package filer

import (
	"fmt"
	"github.com/redhoe/couress/utils/simple"
	"io"
	"mime/multipart"
	"os"
)

var localEngine *LocalOss

type LocalOss struct {
}

func LocalEngin() *LocalOss {
	if localEngine == nil {
		localEngine = &LocalOss{}
	}
	return localEngine
}

// UploadMultipart 文件上传(文件流) saveLocalPath:本地保存绝对路径
func (*LocalOss) UploadMultipart(file *multipart.FileHeader, fileName, saveLocalPath, baseUrl string) (string, error) {
	// 打开文件源
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("FileOpenError:%v", err)
	}
	defer src.Close()
	sPath := simple.ToPath(saveLocalPath, fileName) // 本地保存绝对路径
	visitPath := simple.ToPath(baseUrl, fileName)   // web页面访问Url
	dst, err := os.Create(sPath)
	if err != nil {
		return "", fmt.Errorf("FileSaveError:%v", err)
	}
	defer dst.Close()
	// 将源拷贝到目标文件
	if _, err = io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("FileCopyError:%v", err)
	}
	return visitPath, nil
}

// UploadReader 文件上传(io.Reader)
func (*LocalOss) UploadReader(src io.Reader, fileName, saveLocalPath, baseUrl string) (string, error) {
	// 下面创建保存路径文件 filer.Filename 即上传文件的名字 创建upload文件夹
	sPath := simple.ToPath(saveLocalPath, fileName) // 本地保存绝对路径
	visitPath := simple.ToPath(baseUrl, fileName)   // web页面访问Url
	dst, err := os.Create(sPath)
	if err != nil {
		return "", fmt.Errorf("FileSaveError:%v", err)
	}
	defer dst.Close()
	// 下面将源拷贝到目标文件
	if _, err = io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("FileCopyError:%v", err)
	}
	return visitPath, nil
}
