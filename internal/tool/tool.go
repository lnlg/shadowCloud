package tool

import (
	"os"
	"path/filepath"
)

// 获取当前执行文件绝对路径（go run 和 go build 都可以获取到）
func GetRootDir() string {
	rootPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return rootPath
}

// 获取当前执行文件绝对路径
func GetCurrentAbPathByExecutable() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	realPath, err := filepath.EvalSymlinks(filepath.Dir(ex))
	if err != nil {
		panic(err)
	}
	return filepath.Dir(realPath)
}
