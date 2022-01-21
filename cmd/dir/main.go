package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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

	srcPath, err := handler.initSrcPath()
	if err != nil {
		return
	}

	fi, err := handler.getFileInfos(srcPath)
	if err != nil {
		return
	}

	// 读取文件信息
	descFi, err := handler.getFileInfos(srcPath)
	if err != nil {
		return
	}

	// 原名称
	var originNameSlice = make([]string, len(fi))
	for i := range fi {
		originNameSlice[i] = fi[i].Name()
	}

	ossort.FileInfoAsc(fi)
	ossort.FileInfoDesc(descFi)

	fmt.Printf("%-36s %-36s %-36s\n", "origin", "sorted-asc", "sorted-desc")
	for i := range originNameSlice {
		fmt.Printf("%-36s %-36s %-36s\n", originNameSlice[i], fi[i].Name(), descFi[i].Name())
	}
}

type test struct{}

// getFileInfos .
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

func (s *test) initSrcPath() (srcPath string, err error) {
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
