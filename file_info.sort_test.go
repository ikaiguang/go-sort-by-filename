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

	// 读取文件信息
	fi, err := readDir(srcPath)
	require.Nil(t, err)

	// 读取文件信息
	descFi, err := readDir(srcPath)
	require.Nil(t, err)

	t.Log(len(fi))
	t.Log(len(descFi))
	t.Log(len(fi) == len(descFi))

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
