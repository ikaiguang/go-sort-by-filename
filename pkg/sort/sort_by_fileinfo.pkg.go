package sortpkg

import (
	"os"
)

// FileInfo 文件信息
type FileInfo struct {
	os.FileInfo
	identifier string
}

// SortIdentifier 排序标识
func (s *FileInfo) SortIdentifier() string {
	return s.identifier
}

// Value .
func (s *FileInfo) Value() os.FileInfo {
	return s.FileInfo
}

// NewFileInfo .
func NewFileInfo(fi os.FileInfo) *FileInfo {
	return &FileInfo{
		FileInfo:   fi,
		identifier: GenSortIdentifier(fi.Name()),
	}
}
