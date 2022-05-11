package ossort

import (
	"sort"
)

// filename 按文件名排序
type filename struct{}

// NewSortByFilename .
func NewSortByFilename() SortByFilename {
	return &filename{}
}

// Identifier 排序凭证
func (s *filename) Identifier(nameSlice []string) (newNameSlice []*SortIdentifier) {
	newNameSlice = make([]*SortIdentifier, len(nameSlice))
	for i := range nameSlice {
		newNameSlice[i] = &SortIdentifier{
			OriginData: nameSlice[i],
			SortKey:    GenSortName(nameSlice[i]),
		}
	}
	return newNameSlice
}

// Asc 文件名升序排序
func (s *filename) Asc(nameSlice []string) {
	sortSlice := s.Identifier(nameSlice)
	sort.Slice(sortSlice, func(i, j int) bool {
		return sortSlice[i].SortKey < sortSlice[j].SortKey
	})
	for i := range sortSlice {
		nameSlice[i] = sortSlice[i].OriginData.(string)
	}
	return
}

// Desc 文件名倒序排序
func (s *filename) Desc(nameSlice []string) {
	sortSlice := s.Identifier(nameSlice)
	sort.Slice(sortSlice, func(i, j int) bool {
		return sortSlice[i].SortKey > sortSlice[j].SortKey
	})
	for i := range sortSlice {
		nameSlice[i] = sortSlice[i].OriginData.(string)
	}
	return
}

// Slice 自定义排序 sort.Sort
func (s *filename) Slice(nameSlice []string, less func(i, j int) bool) {
	sort.Slice(nameSlice, less)
}

// SliceStable 自定义排序 sort.SliceStable
func (s *filename) SliceStable(nameSlice []string, less func(i, j int) bool) {
	sort.SliceStable(nameSlice, less)
}
