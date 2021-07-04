package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func main() {
	/*
	整型
	整型中主要有两大类：
	按照整型长度划分：int8, int16, int32, int64
	按照有无符号划分：uint8, uint16, uint32, uint64
	整数类型之间可以互相转换，高长度类型向低长度类型转换会发生长度截取，仅会保留高长度类型的低位值，造成转换错误
	*/
	var a uint16 = math.MaxUint8 + 1
	fmt.Printf("a valud(uint16) is %v\n", a)
	var b = uint8(a)
	fmt.Printf("b valud(uint8) is %v\n", b)


	/*
	浮点型
	浮点型主要有两种：
	float32
	float64
	float32和float64之间也可以进行类型转换，但仍要注意转换之间精度的损失
	*/
	fmt.Printf("%f\n", math.E)
	fmt.Printf("%.2f\n", math.E)

	/*
	布尔型
	Go中的布尔型即常见的true和false。Go中布尔型无法转成整型，无法进行计算
	*/


	/*
	字符串型
	字符串有以下两种方式：
	byte，类型为"uint8"，代表一个ASCII字符
	rune，类型为"int32"，代表一个UTF8字符
	*/
	f := "Golang编程"
	fmt.Printf("byte len f is %v\n", len(f))
	fmt.Printf("rune len f is %v\n", utf8.RuneCountInString(f))
	//按字节遍历字符串
	for _, r := range []byte(f) {
		fmt.Printf("%c", r)
	}
	fmt.Printf("\n")
	//按字符遍历字符串
	for _, r := range f {
		fmt.Printf("%c", r)
	}
	fmt.Printf("\n")
}