package sortutil

import (
	"fmt"
	"strings"
	"testing"
)

// go test -v -count=1 ./ -test.run=TestSort_String
func TestSort_String(t *testing.T) {
	originSlice := []string{"11a.txt", "11a11.txt", "123.txt", "1a.txt", "1a1.txt", "2a.txt", "2a2.txt", "A1.txt", "B1x.txt", "a-1.txt", "a-11.txt", "a-2.txt", "a11x.txt", "a1x.txt", "a2x.txt", "b1.txt", "中国.txt", "微信.txt"}
	var (
		ascSlice  = make([]string, len(originSlice))
		descSlice = make([]string, len(originSlice))
	)
	copy(ascSlice, originSlice)
	copy(descSlice, originSlice)

	sortAscSlice := StringAcs(ascSlice)
	sortDescSlice := StringDesc(descSlice)

	tableSep := strings.Repeat("-", 36)
	fmt.Printf("| %-36s | %-36s | %-36s |\n", "排序：程序编码", "排序：升序", "排序：降序")
	fmt.Printf("| %s | %s | %s |\n", tableSep, tableSep, tableSep)
	for i := range originSlice {
		fmt.Printf("| %-36s | %-36s | %-36s |\n", originSlice[i], sortAscSlice[i].Value(), sortDescSlice[i].Value())
	}
}
