package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	ossort "github.com/ikaiguang/go-sort-by-filename"
)

// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/ ./cmd/dir/...
// CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/ ./cmd/dir/...
// CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/ ./cmd/dir/...
func main() {
	handler := &test{}
	var err error
	defer func() {
		if err != nil {
			log.Println("==> 发送错误", err.Error())
			fmt.Println("===== stack start =====")
			fmt.Printf("%+v\n", err)
			fmt.Println("===== stack end =====")
			os.Exit(1)
		}
	}()

	// 文件目录
	srcPath, err := handler.validateSrcPath()
	if err != nil {
		return
	}

	// 读取文件信息
	ascFi, err := handler.getFileInfos(srcPath)
	if err != nil {
		return
	}
	// 再次读取文件信息
	descFi, err := handler.getFileInfos(srcPath)
	if err != nil {
		return
	}

	// 原名称：程序编码排序
	var originNameSlice = make([]string, len(ascFi))
	for i := range ascFi {
		originNameSlice[i] = ascFi[i].Name()
	}

	// 顺序排序
	ossort.FileInfoAsc(ascFi)
	// 倒叙排序
	ossort.FileInfoDesc(descFi)

	tableSep := strings.Repeat("-", 36)
	fmt.Printf("| %-36s | %-36s | %-36s |\n", "排序：程序编码", "-排序：升序", "排序：降序")
	fmt.Printf("| %s | %s | %s |\n", tableSep, tableSep, tableSep)
	for i := range originNameSlice {
		fmt.Printf("| %-36s | %-36s | %-36s |\n", originNameSlice[i], ascFi[i].Name(), descFi[i].Name())
	}
}

type test struct{}

// getFileInfos 获取目录文件信息
func (s *test) getFileInfos(srcPath string) (fi []os.FileInfo, err error) {
	// 遍历所有的目录
	//walkDirFn := func(path string, d fs.DirEntry, err error) error {
	//	if d.IsDir() {
	//		return nil
	//	}
	//	filenameSlice = append(filenameSlice, d.Name())
	//	return nil
	//}
	//err = filepath.WalkDir(srcPath, walkDirFn)
	//if err != nil {
	//	err = errors.WithStack(err)
	//	return
	//}

	// 当前目录
	d, err := os.Open(srcPath)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	defer func() { _ = d.Close() }()

	fi, err = d.Readdir(-1)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

// validateSrcPath 验证源目录
func (s *test) validateSrcPath() (srcPath string, err error) {
	srcPath = "."
	if len(os.Args) > 1 {
		srcPath = os.Args[1]
	}
	info, err := os.Stat(srcPath)
	if err != nil {
		err = errors.New("无效的文件目录")
		return srcPath, err
	}
	if !info.IsDir() {
		srcPath = filepath.Dir(srcPath)
	}
	return srcPath, err
}
