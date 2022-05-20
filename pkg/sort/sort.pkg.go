package sortpkg

import (
	"sort"
)

// Sortable 可排序
// 泛型文档地址：https://go.dev/blog/when-generics
type Sortable interface {
	SortIdentifier() string
}

// SortAsc 升序排序
func SortAsc[T Sortable](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i].SortIdentifier() < s[j].SortIdentifier()
	})
}

// SortDesc 降序排序
func SortDesc[T Sortable](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i].SortIdentifier() > s[j].SortIdentifier()
	})
}
