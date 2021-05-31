# Go语言数组

## 数组

数组是具有相同数据类型的数据项组成的一组长度固定的序列，数据项叫做数组的元素，数
组的长度必须是非负整数的常量，长度也是类型的一部分

位置: runtime/slice.go

```go
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
```

初始化逻辑: makeslice

```go
func makeslice(et *_type, len, cap int) unsafe.Pointer {
    ...
}
```

1 声明
数组声明需要指定组成元素的类型以及存储元素的数量（长度）。在数组声明后，其长度不
可修改，数组的每个元素会根据对应类型的零值对进行初始化

2 字面量
a) 指定数组长度: [length]type{v1, v2, …, vlength}
b) 使用初始化元素数量推到数组长度: […]type{v1, v2, …, vlength}
c) 对指定位置元素进行初始化: [length]type{im:vm, …, sin:in}

3 操作

+ 关系运算==、!=
+ 获取数组长度 使用 len 函数可获取数组的长度
+ 访问&修改 通过对编号对数组元素进行访问和修改，元素的编号从左到右依次为:0, 1, 2, …, n(n为数组长度-1)
+ 切片: array\[start:end\]获取数组的一部分元素做为切片
+ 遍历 可以通过 for+len+访问方式或 for-range 方式对数组中元素进行遍历使用 for-range 遍历数组，range 返回两个元素分别为数组元素索引和值

4 多维数组

数组的元素也可以是数组类型，此时称为多维数组

+ 声明&初始化
+ 访问&修改
+ 遍历