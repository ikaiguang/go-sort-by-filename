package dirpkg

import (
	"os"
)

// ReadDir .
func ReadDir(srcPath string) (fi []os.FileInfo, err error) {
	// 处理目录
	d, err := os.Open(srcPath)
	if err != nil {
		return fi, err
	}
	defer func() { _ = d.Close() }()

	// 读取文件信息
	fi, err = d.Readdir(-1)
	if err != nil {
		return fi, err
	}
	return fi, err
}
