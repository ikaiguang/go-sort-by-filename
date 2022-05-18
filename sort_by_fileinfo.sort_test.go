package ossort

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v -count=1 ./ -test.run=TestFileInfo_Sort
func TestFileInfo_Sort(t *testing.T) {
	srcPath := "./testdata"
	//srcPath := "./test/fish_gif"

	// 读取文件信息
	fi, err := ReadDir(srcPath)
	require.Nil(t, err)

	// 读取文件信息
	descFi, err := ReadDir(srcPath)
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

	tableSep := strings.Repeat("-", 36)
	fmt.Printf("| %-36s | %-36s | %-36s |\n", "排序：程序编码", "排序：升序", "排序：降序")
	fmt.Printf("| %s | %s | %s |\n", tableSep, tableSep, tableSep)
	for i := range originNameSlice {
		fmt.Printf("| %-36s | %-36s | %-36s |\n", originNameSlice[i], fi[i].Name(), descFi[i].Name())
	}
	//sort.Strings(originNameSlice)
	//for i := range originNameSlice {
	//	preInt64Str, middleStr, sufInt64Str, ext := SplitFilename(originNameSlice[i])
	//	fmt.Println(originNameSlice[i], preInt64Str, middleStr, sufInt64Str, ext)
	//}
}
