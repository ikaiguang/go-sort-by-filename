package ossort

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v -count=1 ./ -test.run=TestFileInfo_Sort
func TestFileInfo_Sort(t *testing.T) {
	srcPath := "./testdata"
	//srcPath := "./test/fish_gif"

	// 读取文件信息
	fi, err := readDir(srcPath)
	require.Nil(t, err)

	// 读取文件信息
	descFi, err := readDir(srcPath)
	require.Nil(t, err)

	// 原名称
	var originNameSlice = make([]string, len(fi))
	for i := range fi {
		originNameSlice[i] = fi[i].Name()
	}

	// 排序后
	handler := NewSortByFileInfo()
	handler.Asc(fi)
	handler.Desc(descFi)

	fmt.Printf("%-36s %-36s %-36s\n", "origin", "sorted-asc", "sorted-desc")
	for i := range originNameSlice {
		fmt.Printf("%-36s %-36s %-36s\n", originNameSlice[i], fi[i].Name(), descFi[i].Name())
	}
	//sort.Strings(originNameSlice)
	//for i := range originNameSlice {
	//	preInt64Str, middleStr, sufInt64Str, ext := SplitFilename(originNameSlice[i])
	//	fmt.Println(originNameSlice[i], preInt64Str, middleStr, sufInt64Str, ext)
	//}
}

// readDir .
func readDir(srcPath string) (fi []os.FileInfo, err error) {
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
