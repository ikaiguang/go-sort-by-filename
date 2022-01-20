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