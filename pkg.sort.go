package ossort

import (
	"encoding/binary"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

var (
	// gbkEncoder gbk encoder
	gbkEncoder = simplifiedchinese.GB18030.NewEncoder()
)

// GenSortName 排序的名称
// 分割为：前面数字 中间文字 后面数字 文件后缀
func GenSortName(filename string) string {
	preInt64Str, middleStr, sufInt64Str, ext := SplitFilename(filename)

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
	if gbkBytes, gbkErr := gbkEncoder.Bytes([]byte(middleStr)); gbkErr == nil {
		middleStr = string(gbkBytes)
	}

	// prefix + numeric-suffix + ext
	return prefixB64Str + middleStr + suffixB64Str + ext
}

// SplitFilename 切分文件名
func SplitFilename(filename string) (preInt64Str, middleStr, sufInt64Str, ext string) {
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
		sufInt64Str = string(name[i]) + sufInt64Str
	}

	if len(preInt64Str) != len(name) {
		middleStr = name[len(preInt64Str) : len(name)-len(sufInt64Str)]
	}
	return
}
