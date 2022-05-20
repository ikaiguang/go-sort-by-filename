package sortutil

import (
	"fmt"
	"strings"
	"testing"

	dirpkg "github.com/ikaiguang/go-sort-by-filename/v2/pkg/dir"
	"github.com/stretchr/testify/require"
)

// go test -v -count=1 ./ -test.run=TestSort_FileInfo
func TestSort_FileInfo(t *testing.T) {
	srcPath := "./testdata"
	//srcPath := "./test/fish_gif"

	// 读取文件信息
	ascFi, err := dirpkg.ReadDir(srcPath)
	require.Nil(t, err)

	// 读取文件信息
	descFi, err := dirpkg.ReadDir(srcPath)
	require.Nil(t, err)

	// 原名称
	var (
		originNameSlice = make([]string, len(ascFi))
	)
	for i := range ascFi {
		originNameSlice[i] = ascFi[i].Name()
	}

	// 排序后
	ascFiSort := FileInfoAsc(ascFi)
	descFiSort := FileInfoDesc(descFi)

	tableSep := strings.Repeat("-", 36)
	fmt.Printf("| %-36s | %-36s | %-36s |\n", "排序：程序编码", "排序：升序", "排序：降序")
	fmt.Printf("| %s | %s | %s |\n", tableSep, tableSep, tableSep)
	for i := range originNameSlice {
		fmt.Printf("| %-36s | %-36s | %-36s |\n", originNameSlice[i], ascFiSort[i].Value().Name(), descFiSort[i].Value().Name())
	}
	//sort.Strings(originNameSlice)
	//for i := range originNameSlice {
	//	preInt64Str, middleStr, sufInt64Str, ext := SplitFilename(originNameSlice[i])
	//	fmt.Println(originNameSlice[i], preInt64Str, middleStr, sufInt64Str, ext)
	//}
}
