package sortutil

import (
	"sort"

	sortpkg "github.com/ikaiguang/go-sort-by-filename/v2/pkg/sort"
)

// Asc 升序排序
func Asc[T sortpkg.Sortable](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i].SortIdentifier() < s[j].SortIdentifier()
	})
}

// Desc 降序排序
func Desc[T sortpkg.Sortable](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i].SortIdentifier() > s[j].SortIdentifier()
	})
}

// SortAsc 升序排序
func SortAsc[T sortpkg.Sortable](s []T) {
	sortpkg.SortAsc(s)
}

// SortDesc 降序排序
func SortDesc[T sortpkg.Sortable](s []T) {
	sortpkg.SortDesc(s)
}
