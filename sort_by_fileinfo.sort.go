package ossort

import (
	"os"
	"sort"
)

// fileInfo 按文件名排序
type fileInfo struct{}

// NewSortByFileInfo .
func NewSortByFileInfo() SortByFileInfo {
	return &fileInfo{}
}

// Identifier 排序凭证
func (s *fileInfo) Identifier(fi []os.FileInfo) (newInfoSlice []*SortIdentifier) {
	newInfoSlice = make([]*SortIdentifier, len(fi))
	for i := range fi {
		newInfoSlice[i] = &SortIdentifier{
			OriginData: fi[i],
			SortKey:    GenSortName(fi[i].Name()),
		}
	}
	return newInfoSlice
}

// Asc 文件名升序排序
func (s *fileInfo) Asc(fi []os.FileInfo) {
	sortSlice := s.Identifier(fi)
	sort.Slice(sortSlice, func(i, j int) bool {
		return sortSlice[i].SortKey < sortSlice[j].SortKey
	})
	for i := range sortSlice {
		fi[i] = sortSlice[i].OriginData.(os.FileInfo)
	}
	return
}

// Desc 文件名倒序排序
func (s *fileInfo) Desc(fi []os.FileInfo) {
	sortSlice := s.Identifier(fi)
	sort.Slice(sortSlice, func(i, j int) bool {
		return sortSlice[i].SortKey > sortSlice[j].SortKey
	})
	for i := range sortSlice {
		fi[i] = sortSlice[i].OriginData.(os.FileInfo)
	}
	return
}

// Slice 自定义排序 sort.Sort
func (s *fileInfo) Slice(fi []os.FileInfo, less func(i, j int) bool) {
	sort.Slice(fi, less)
}

// SliceStable 自定义排序 sort.SliceStable
func (s *fileInfo) SliceStable(fi []os.FileInfo, less func(i, j int) bool) {
	sort.SliceStable(fi, less)
}
