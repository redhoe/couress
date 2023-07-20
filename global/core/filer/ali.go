package filer

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/redhoe/couress/global/core/confer"
	"github.com/redhoe/couress/utils/simple"
	"io"
	"mime/multipart"
)

var aliEngine *AliOss

type AliOss struct {
	ossClient  *oss.Client
	bucketName string
	bucketUrl  string // 访问地址
	basePath   string //上传路径
}

func AliEngin() *AliOss {
	if aliEngine == nil {
		client, err := oss.New(
			confer.AppConfServer.Oss.Endpoint,
			confer.AppConfServer.Oss.AccessKeyId,
			confer.AppConfServer.Oss.AccessKeySecret)
		if err != nil {
			panic(err)
		}
		err = client.SetBucketACL(confer.AppConfServer.Oss.BucketName, oss.ACLPublicReadWrite)
		if err != nil {
			panic(err)
		}
		aliEngine = &AliOss{
			client,
			confer.AppConfServer.Oss.BucketName,
			confer.AppConfServer.Oss.BucketUrl,
			confer.AppConfServer.Oss.BasePath}
	}
	return aliEngine
}

// UploadMultipart 文件上传(文件流)
func (s *AliOss) UploadMultipart(file *multipart.FileHeader, fileName, savePath string) (string, error) {
	// 打开文件源
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("FileOpenError:%v", err)
	}
	defer src.Close()

	bucket, err := s.ossClient.Bucket(s.bucketName)
	if err != nil {
		return "", fmt.Errorf("OssBucketError:%v", err)
	}
	// 上传文件

	objectName := simple.ToPath(s.basePath, savePath, fileName)             // 本地保存绝对路径
	visitPath := simple.ToPath(s.bucketUrl, s.basePath, savePath, fileName) // web页面访问Url

	// 将文件流上传至upload下

	err = bucket.PutObject(objectName, src)
	if err != nil {
		fmt.Println("Error:", err)
		return "", fmt.Errorf("PutObject:%v", err)
	}
	return visitPath, nil

}

// UploadReader 文件上传(io.Reader)
func (s *AliOss) UploadReader(src io.Reader, fileName, savePath string) (string, error) {
	// 下面创建保存路径文件 filer.Filename 即上传文件的名字 创建upload文件夹
	objectName := simple.ToPath(s.basePath, savePath, fileName) // 本地保存绝对路径
	visitPath := simple.ToPath(s.bucketUrl, savePath, fileName) // web页面访问Url
	bucket, err := s.ossClient.Bucket(s.bucketName)
	if err != nil {
		return "", fmt.Errorf("OssBucketError:%v", err)
	}
	err = bucket.PutObject(objectName, src)
	if err != nil {
		fmt.Println("Error:", err)
		return "", fmt.Errorf("OssPutObject:%v", err)
	}
	return visitPath, nil
}

// UploadFileToOss  文件上传
func (s *AliOss) UploadFileToOss(localFilePath, fileName, savePath string) (string, error) {
	bucket, err := s.ossClient.Bucket(s.bucketName)
	if err != nil {
		return "", err
	}
	objectName := simple.ToPath(s.basePath, savePath, fileName) // 本地保存绝对路径
	visitPath := simple.ToPath(s.bucketUrl, savePath, fileName) // web页面访问Url
	err = bucket.PutObjectFromFile(objectName, localFilePath)
	if err != nil {
		return "", err
	}
	return visitPath, nil
}
