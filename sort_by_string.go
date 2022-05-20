package sortutil

import (
	sortpkg "github.com/ikaiguang/go-sort-by-filename/v2/pkg/sort"
)

// StringAcs 升序排序
func StringAcs(s []string) []*sortpkg.String {
	var sortSlice = make([]*sortpkg.String, len(s))
	for i := range s {
		sortSlice[i] = sortpkg.NewString(s[i])
	}
	sortpkg.SortAsc(sortSlice)
	return sortSlice
}

// StringDesc 降序排序
func StringDesc(s []string) []*sortpkg.String {
	var sortSlice = make([]*sortpkg.String, len(s))
	for i := range s {
		sortSlice[i] = sortpkg.NewString(s[i])
	}
	sortpkg.SortDesc(sortSlice)
	return sortSlice
}
