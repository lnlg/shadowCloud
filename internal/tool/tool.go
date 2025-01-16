package tool

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"time"
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

// 返回当前时间：2006-01-02 15:04:05
func GetNowDate() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}

// 返回当前时间戳: 1580403200
func GetDateTime() int {
	return int(time.Now().Unix())
}

// 加密MD5
func EncryMd5(s string) string {
	ctx := md5.New() // 初始化一个MD5对象
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}

// json解码
func JsonDecode(s string) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 扫描指定目录
func ScanDirectory(dirPath string) []string {
	var files []string
	err := filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil
	}
	return files
}

// 判断文件是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
