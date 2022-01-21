# golang sorted by filename

os sorted by file name

**Note** : 文件名称包涵特殊字符(-_)，排序的位置根据视操作系统语言而定

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

origin                               sorted-asc                           sorted-desc                         
B1x.txt                              1a.txt                               中国.txt                              
中国.txt                              1a1.txt                              微信.txt                              
2a2.txt                              2a.txt                               special                             
11a.txt                              2a2.txt                              B1x.txt                             
1a.txt                               11a.txt                              b1.txt                              
a-2.txt                              11a11.txt                            a2x.txt                             
a11x.txt                             123.txt                              a1x.txt                             
a-1.txt                              A1.txt                               a11x.txt                            
微信.txt                              a-1.txt                              a-11.txt                            
A1.txt                               a-2.txt                              a-2.txt                             
1a1.txt                              a-11.txt                             a-1.txt                             
a1x.txt                              a11x.txt                             A1.txt                              
special                              a1x.txt                              123.txt                             
a-11.txt                             a2x.txt                              11a11.txt                           
11a11.txt                            b1.txt                               11a.txt                             
b1.txt                               B1x.txt                              2a2.txt                             
a2x.txt                              special                              2a.txt                              
123.txt                              微信.txt                              1a1.txt                             
2a.txt                               中国.txt                              1a.txt 

```

## Test

```shell

go test -v -count=1 ./ -test.run=TestFileInfo_Sort

go test -v -count=1 ./ -test.run=TestFilename_Sort

```