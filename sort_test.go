package sortutil

import (
	"os"

	dirpkg "github.com/ikaiguang/go-sort-by-filename/v2/pkg/dir"
	sortpkg "github.com/ikaiguang/go-sort-by-filename/v2/pkg/sort"
)

// File 文件：例子
type File struct {
	os.FileInfo
	sortIdentifier string
}

// SortIdentifier .
func (s *File) SortIdentifier() string {
	return s.sortIdentifier
}

// ExampleAsc 升序排序
func ExampleAsc() {
	srcPath := "./testdata"
	// 读取文件信息
	fi, err := dirpkg.ReadDir(srcPath)
	if err != nil {
		return
	}
	var fileSlice = make([]*File, len(fi))
	for i := range fi {
		fileSlice[i] = &File{FileInfo: fi[i], sortIdentifier: sortpkg.GenSortIdentifier(fi[i].Name())}
	}
	Asc(fileSlice)
	// 或者 sortpkg.SortAsc(s)
}

// ExampleDesc 降序排序
func ExampleDesc() {
	srcPath := "./testdata"
	// 读取文件信息
	fi, err := dirpkg.ReadDir(srcPath)
	if err != nil {
		return
	}
	var fileSlice = make([]*File, len(fi))
	for i := range fi {
		fileSlice[i] = &File{FileInfo: fi[i], sortIdentifier: sortpkg.GenSortIdentifier(fi[i].Name())}
	}
	Desc(fileSlice)
	// 或者 sortpkg.SortDesc(s)
}
