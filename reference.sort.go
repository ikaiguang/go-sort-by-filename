package ossort

// SortInterface 排序接口
type SortInterface interface {
	Asc(nameSlice []string)
	Desc(nameSlice []string)
}
