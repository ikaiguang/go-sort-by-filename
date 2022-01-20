package ossort

import (
	"encoding/binary"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// filename 按文件名排序
type filename struct{}

// NewSortByFilename .
func NewSortByFilename() SortInterface {
	return &filename{}
}

// SortByFilename .
type SortByFilename struct {
	OriginName string
	SortName   string
}

// Asc 升序排序
func (s *filename) Asc(nameSlice []string) {
	sortSlice := s.parseFilename(nameSlice)
	sort.Slice(sortSlice, func(i, j int) bool {
		return sortSlice[i].SortName < sortSlice[j].SortName
	})
	for i := range sortSlice {
		nameSlice[i] = sortSlice[i].OriginName
	}
	return
}

// Desc 倒序排序
func (s *filename) Desc(nameSlice []string) {
	sortSlice := s.parseFilename(nameSlice)
	sort.Slice(sortSlice, func(i, j int) bool {
		return sortSlice[i].SortName > sortSlice[j].SortName
	})
	for i := range sortSlice {
		nameSlice[i] = sortSlice[i].OriginName
	}
	return
}

// parseFilename 解析文件名
func (s *filename) parseFilename(nameSlice []string) (newNameSlice []*SortByFilename) {
	// 转码 gbk
	encoder := simplifiedchinese.GBK.NewEncoder()

	newNameSlice = make([]*SortByFilename, len(nameSlice))
	for i := range nameSlice {
		newNameSlice[i] = &SortByFilename{
			OriginName: nameSlice[i],
			SortName:   s.genSortName(encoder, nameSlice[i]),
		}
	}
	return newNameSlice
}

// genSortName 排序的名称
// 分割为：前面数字 中间文字 后面数字 文件后缀
func (s *filename) genSortName(encoder *encoding.Encoder, filename string) string {
	preInt64Str, middleStr, sufInt64Str, ext := s.splitFilename(filename)

	// string numeric prefix to uint64 bytes
	// empty string is zero, so integers are plus one
	prefixB64 := make([]byte, 64/8)
	prefixB64Str := ""
	if len(preInt64Str) > 0 {
		u64, err := strconv.ParseUint(preInt64Str, 10, 64)
		if err == nil {
			binary.BigEndian.PutUint64(prefixB64, u64+1)
		}
		prefixB64Str = string(prefixB64)
	} else {
		prefixB64Str = "a"
	}

	// string numeric suffix to uint64 bytes
	// empty string is zero, so integers are plus one
	suffixB64 := make([]byte, 64/8)
	suffixB64Str := ""
	if len(sufInt64Str) > 0 {
		u64, err := strconv.ParseUint(sufInt64Str, 10, 64)
		if err == nil {
			binary.BigEndian.PutUint64(suffixB64, u64+1)
		}
		suffixB64Str = string(suffixB64)
	} else {
		binary.BigEndian.PutUint64(suffixB64, 0)
		suffixB64Str = string(suffixB64)
	}

	// middle
	middleStr = strings.ToLower(middleStr)
	if gbkBytes, gbkErr := encoder.Bytes([]byte(middleStr)); gbkErr == nil {
		middleStr = string(gbkBytes)
	}

	// prefix + numeric-suffix + ext
	return prefixB64Str + middleStr + suffixB64Str + ext
}

// splitFilename 切分文件名
func (s *filename) splitFilename(filename string) (preInt64Str, middleStr, sufInt64Str, ext string) {
	ext = filepath.Ext(filename)
	name := filename[:len(filename)-len(ext)]

	// split numeric prefix
	maxNameIndex := len(name) - 1
	for i := 0; i <= maxNameIndex; i++ {
		if '0' > name[i] || name[i] > '9' {
			break
		}
		preInt64Str += string(name[i])
	}

	// split numeric suffix
	for i := maxNameIndex; i >= 0; i-- {
		if '0' > name[i] || name[i] > '9' {
			break
		}
		sufInt64Str += string(name[i])
	}

	if len(preInt64Str) != len(name) {
		middleStr = name[len(preInt64Str) : len(name)-len(sufInt64Str)]
	}
	return
}
