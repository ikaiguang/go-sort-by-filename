# golang sorted by filename

os sorted by file name

- 实现 Windows OS 按文件名排序
- 实现 MAC OS 按文件名排序

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

## test

go test -v -count=1 ./ -test.run=TestFilename_Sort

```go
package ossort

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./ -test.run=TestFilename_Sort
func TestFilename_Sort(t *testing.T) {
	nameSlice1 := []string{"11a.txt", "11a11.txt", "123.txt", "1a.txt", "1a1.txt", "2a.txt", "2a2.txt", "A1.txt", "B1x.txt", "a-1.txt", "a-11.txt", "a-2.txt", "a11x.txt", "a1x.txt", "a2x.txt", "b1.txt", "中国.txt", "微信.txt"}
	nameSlice2 := []string{"11a.txt", "11a11.txt", "123.txt", "1a.txt", "1a1.txt", "2a.txt", "2a2.txt", "A1.txt", "B1x.txt", "a-1.txt", "a-11.txt", "a-2.txt", "a11x.txt", "a1x.txt", "a2x.txt", "b1.txt", "中国.txt", "微信.txt"}
	tests := []*struct {
		name      string
		direction string
		nameSlice []string
		want      string
	}{
		{
			name:      "#Sort_By_Filename_#ASC",
			direction: "asc",
			nameSlice: nameSlice1,
			want:      "1a.txt,1a1.txt,2a.txt,2a2.txt,11a.txt,11a11.txt,123.txt,A1.txt,a-1.txt,a-2.txt,a-11.txt,a11x.txt,a1x.txt,a2x.txt,b1.txt,B1x.txt,微信.txt,中国.txt",
		},
		{
			name:      "#Sort_By_Filename_#DESC",
			direction: "desc",
			nameSlice: nameSlice2,
			want:      "中国.txt,微信.txt,B1x.txt,b1.txt,a2x.txt,a1x.txt,a11x.txt,a-11.txt,a-2.txt,a-1.txt,A1.txt,123.txt,11a11.txt,11a.txt,2a2.txt,2a.txt,1a1.txt,1a.txt",
		},
	}

	handler := NewSortByFilename()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch strings.ToLower(tt.direction) {
			case "asc":
				handler.Asc(tt.nameSlice)
			default:
				handler.Desc(tt.nameSlice)
			}
			got := strings.Join(tt.nameSlice, ",")
			if tt.want != got {
				assert.Equal(t, tt.want, got)
				t.Fail()
			}
		})
	}
}

```