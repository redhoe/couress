package simple

import (
	"os"
	"strings"
	"testing"
)

func TestPathAdd(t *testing.T) {
	a := "/a/a1/a2/"
	b := "/b"
	c := "c/"
	d := " a b c /s "
	t.Log("1", ToPath(a, b, c, d), "end")
	t.Log("2", strings.TrimSuffix(a, "/"), "end")
	t.Log("3", strings.Trim(a, "/"), "end")
	t.Log("4", strings.TrimLeft(a, "/"), "end")
	t.Log("5", strings.TrimRight(a, "/"), "end")
	t.Log("6", strings.TrimSpace(d), "end")

}

func TestPathAdd02(t *testing.T) {
	a := "http://www.baidu.com/"
	b := "/b/"
	c := "c/"
	d := "/s"
	t.Log("1", ToPath(a, b, c, d), "...end")
}

func TestPathMk(t *testing.T) {
	savePath := AbsPath("hello", "fuck")
	t.Log(savePath)
}

// 获取文件信息
func TestOsStat(t *testing.T) {
	file := "./floats.go"
	//file := "./index.png"
	fileInfo, err := os.Stat(file)
	t.Log(err)
	t.Log("Size", fileInfo.Size())
	t.Log("Name", fileInfo.Name())
	t.Log("IsDir", fileInfo.IsDir())
	t.Log("Mode", fileInfo.Mode())
	t.Log("Mode", fileInfo.Mode().String())
	t.Log("Mode", fileInfo.Mode().IsDir())
	t.Log("Mode IsRegular", fileInfo.Mode().IsRegular())
	t.Log("Mode", fileInfo.Mode().Perm())
	t.Log("Mode", fileInfo.Mode().Type())
	t.Log("ModTime", fileInfo.ModTime())
}
