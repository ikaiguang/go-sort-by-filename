package sortpkg

// String 字符串
type String struct {
	value      string
	identifier string
}

// SortIdentifier 排序标识
func (s *String) SortIdentifier() string {
	return s.identifier
}

// Value .
func (s *String) Value() string {
	return s.value
}

// NewString .
func NewString(s string) *String {
	return &String{
		value:      s,
		identifier: GenSortIdentifier(s),
	}
}
