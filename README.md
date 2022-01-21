# golang sorted by filename

os sorted by file name

**Note** : 文件名称包涵特殊字符(-_)，排序的位置根据视操作系统语言而定

go version go1.16+ 或

- 实现 Windows OS 按文件名排序
- 实现 MAC OS 按文件名排序

## 简单使用

获取包 `go get github.com/ikaiguang/go-sort-by-filename`

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

```txt

    排序前                   排序后
    -1.txt                  1a.txt
    11a.txt                 1a1.txt
    11a11.txt               2a.txt
    123.txt                 2a2.txt
    1a.txt                  11a.txt
    1a1.txt                 11a11.txt
    2a.txt                  123.txt
    2a2.txt                 -1.txt
    A1.txt                  _1.txt
    B1x.txt                 A1.txt
    _1.txt                  a-1.txt
    a-1.txt                 a-2.txt
    a-11.txt                a-11.txt
    a-2.txt                 a11x.txt
    a11x.txt                a1x.txt
    a1x.txt                 a2x.txt
    a2x.txt                 b1.txt
    b1.txt                  B1x.txt
    中国.txt                 微信.txt
    微信.txt                 中国.txt

```

## Test

```shell

go test -v -count=1 ./ -test.run=TestFileInfo_Sort

go test -v -count=1 ./ -test.run=TestFilename_Sort

```