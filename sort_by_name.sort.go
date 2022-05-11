package ossort

import (
	"sort"
)

// Name 排序名称
type Name interface {
	Name() string
}

// sortByName 排序
type sortByName struct{}

// NewSortByName .
func NewSortByName() SortByName {
	return &sortByName{}
}

// Identifier 排序凭证
func (s *sortByName) Identifier(infos []Name) (newInfoSlice []*SortIdentifier) {
	newInfoSlice = make([]*SortIdentifier, len(infos))
	for i := range infos {
		newInfoSlice[i] = &SortIdentifier{
			OriginData: infos[i],
			SortKey:    GenSortName(infos[i].Name()),
		}
	}
	return newInfoSlice
}

// Asc 文件名升序排序
func (s *sortByName) Asc(infos []Name) {
	sortSlice := s.Identifier(infos)
	sort.Slice(sortSlice, func(i, j int) bool {
		return sortSlice[i].SortKey < sortSlice[j].SortKey
	})
	for i := range sortSlice {
		infos[i] = sortSlice[i].OriginData.(Name)
	}
	return
}

// Desc 文件名倒序排序
func (s *sortByName) Desc(infos []Name) {
	sortSlice := s.Identifier(infos)
	sort.Slice(sortSlice, func(i, j int) bool {
		return sortSlice[i].SortKey > sortSlice[j].SortKey
	})
	for i := range sortSlice {
		infos[i] = sortSlice[i].OriginData.(Name)
	}
	return
}

// Slice 自定义排序 sort.Sort
func (s *sortByName) Slice(infos []Name, less func(i, j int) bool) {
	sort.Slice(infos, less)
}

// SliceStable 自定义排序 sort.SliceStable
func (s *sortByName) SliceStable(infos []Name, less func(i, j int) bool) {
	sort.SliceStable(infos, less)
}
