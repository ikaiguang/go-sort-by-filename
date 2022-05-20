package sortutil

import (
	"os"

	sortpkg "github.com/ikaiguang/go-sort-by-filename/v2/pkg/sort"
)

// FileInfoAsc 升序排序
func FileInfoAsc(s []os.FileInfo) []*sortpkg.FileInfo {
	var sortSlice = make([]*sortpkg.FileInfo, len(s))
	for i := range s {
		sortSlice[i] = sortpkg.NewFileInfo(s[i])
	}
	sortpkg.SortAsc(sortSlice)

	return sortSlice
}

// FileInfoDesc 降序排序
func FileInfoDesc(s []os.FileInfo) []*sortpkg.FileInfo {
	var sortSlice = make([]*sortpkg.FileInfo, len(s))
	for i := range s {
		sortSlice[i] = sortpkg.NewFileInfo(s[i])
	}
	sortpkg.SortDesc(sortSlice)

	return sortSlice
}
