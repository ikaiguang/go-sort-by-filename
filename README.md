# golang sorted by filename

os sorted by file name

操作系统中的排序：根据文件名进行升序/降序排序

**Note** : 包涵特殊字符(-_)的文件名，在各个操作系统有一定的差异

## 简单使用

获取SDK `go get github.com/ikaiguang/go-sort-by-filename`

```go

package main

import (
	"os"

	ossort "github.com/ikaiguang/go-sort-by-filename"
)

func main() {
	ossort.FilenameAsc([]string{})
	ossort.FilenameDesc([]string{})
	ossort.FileInfoAsc([]os.FileInfo{})
	ossort.FileInfoDesc([]os.FileInfo{})
}

```

## 排序示例

| 排序：程序编码 | 排序：升序 | 排序：降序 |
| -------------- | ---------- | ---------- |
| B1x.txt        | 1a.txt     | 中国.txt   |
| 中国.txt       | 1a1.txt    | 微信.txt   |
| 2a2.txt        | 2a.txt     | special    |
| 11a.txt        | 2a2.txt    | B1x.txt    |
| 1a.txt         | 11a.txt    | b1.txt     |
| .DS_Store      | 11a11.txt  | a2x.txt    |
| a-2.txt        | 123.txt    | a1x.txt    |
| a11x.txt       | .DS_Store  | a11x.txt   |
| a-1.txt        | A1.txt     | a-11.txt   |
| 微信.txt       | a-1.txt    | a-2.txt    |
| A1.txt         | a-2.txt    | a-1.txt    |
| 1a1.txt        | a-11.txt   | A1.txt     |
| a1x.txt        | a11x.txt   | .DS_Store  |
| special        | a1x.txt    | 123.txt    |
| a-11.txt       | a2x.txt    | 11a11.txt  |
| 11a11.txt      | b1.txt     | 11a.txt    |
| b1.txt         | B1x.txt    | 2a2.txt    |
| a2x.txt        | special    | 2a.txt     |
| 123.txt        | 微信.txt   | 1a1.txt    |
| 2a.txt         | 中国.txt   | 1a.txt     |

## 运行测试用例

```shell

go test -v -count=1 ./ -test.run=TestFileInfo_Sort

go test -v -count=1 ./ -test.run=TestFilename_Sort

```
