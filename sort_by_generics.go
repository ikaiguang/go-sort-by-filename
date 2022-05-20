package sortutil

import (
	sortpkg "github.com/ikaiguang/go-sort-by-filename/v2/pkg/sort"
)

// GenericsAsc 升序排序
func GenericsAsc(s []string) []*sortpkg.AnyGenerics[string] {
	var sortSlice = make([]*sortpkg.AnyGenerics[string], len(s))
	for i := range s {
		sortSlice[i] = sortpkg.NewAnyGenerics(s[i], s[i])
	}
	sortpkg.SortAsc(sortSlice)

	return sortSlice
}

// GenericsDesc 降序排序
func GenericsDesc(s []string) []*sortpkg.AnyGenerics[string] {
	var sortSlice = make([]*sortpkg.AnyGenerics[string], len(s))
	for i := range s {
		sortSlice[i] = sortpkg.NewAnyGenerics(s[i], s[i])
	}
	sortpkg.SortDesc(sortSlice)

	return sortSlice
}
