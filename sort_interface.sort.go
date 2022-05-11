package ossort

import (
	"os"
)

// SortIdentifier .
type SortIdentifier struct {
	OriginData interface{}
	SortKey    string
}

// SortByFilename 排序接口
type SortByFilename interface {
	// Identifier 排序凭证
	Identifier(nameSlice []string) []*SortIdentifier
	// Asc 文件名升序排序
	Asc(nameSlice []string)
	// Desc 文件名倒序排序
	Desc(nameSlice []string)
	// Slice 自定义排序 = sort.Sort
	Slice(nameSlice []string, less func(i, j int) bool)
	// SliceStable 自定义排序 = sort.SliceStable
	SliceStable(nameSlice []string, less func(i, j int) bool)
}

// SortByFileInfo 排序接口
type SortByFileInfo interface {
	// Identifier 排序凭证
	Identifier(fi []os.FileInfo) []*SortIdentifier
	// Asc 文件名升序排序
	Asc(fi []os.FileInfo)
	// Desc 文件名倒序排序
	Desc(fi []os.FileInfo)
	// Slice 自定义排序 = sort.Sort
	Slice(fi []os.FileInfo, less func(i, j int) bool)
	// SliceStable 自定义排序 = sort.SliceStable
	SliceStable(fi []os.FileInfo, less func(i, j int) bool)
}

// SortByName 排序接口
type SortByName interface {
	// Identifier 排序凭证
	Identifier(infos []Name) []*SortIdentifier
	// Asc 文件名升序排序
	Asc(infos []Name)
	// Desc 文件名倒序排序
	Desc(infos []Name)
	// Slice 自定义排序 = sort.Sort
	Slice(infos []Name, less func(i, j int) bool)
	// SliceStable 自定义排序 = sort.SliceStable
	SliceStable(infos []Name, less func(i, j int) bool)
}
