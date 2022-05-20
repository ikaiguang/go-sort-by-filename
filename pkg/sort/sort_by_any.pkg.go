package sortpkg

// AnyGenerics 泛型使用
// Deprecated: 不推荐使用；推荐使用自声明与实现 sortpkg.Sortable
// 参考： sortpkg.String 与 sortpkg.FileInfo
type AnyGenerics[t any] struct {
	value      t
	identifier string
}

// SortIdentifier 排序标识
func (s *AnyGenerics[t]) SortIdentifier() string {
	return s.identifier
}

// Value 原来的值
func (s *AnyGenerics[t]) Value() t {
	return s.value
}

// NewAnyGenerics 创建泛型
// Deprecated: 不推荐使用；推荐使用自声明与实现 sortpkg.Sortable
// 参考： sortpkg.String 与 sortpkg.FileInfo
func NewAnyGenerics[t any](value t, name string) *AnyGenerics[t] {
	return &AnyGenerics[t]{
		value:      value,
		identifier: GenSortIdentifier(name),
	}
}
