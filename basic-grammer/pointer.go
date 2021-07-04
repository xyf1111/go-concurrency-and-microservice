package main

import (
	"flag"
	"fmt"
)

func main() {
	/*
	在Go语言中，指针包含以下三个概念
	指针地址
	指针类型
	指针取值
	*/

	//声明一个string类型
	str := "Golang is Good!"
	//获取str的指针
	strPtr := &str

	/*
	str类型为string，它的指针strPtr的类型为*string，即为变量str在内存中的地址
	*/

	fmt.Printf("str type is %T, and value is %v\n", str, str)
	fmt.Printf("strPtr type is %T, and value is %v\n", strPtr, strPtr)

	strPtrPtr := &strPtr
	fmt.Printf("strPtr type is %T, and value is %v\n", strPtrPtr, strPtrPtr)

	/*
	Go提供根据指针获取变量值的取值操作(*)，通过取值操作可以获取指针对应变量的值和变量进行赋值操作
	*/
	//获取指针对应变量的值
	newStr := *strPtr
	fmt.Printf("newStr type is %T, value is %v, address is %v\n", newStr, newStr, &newStr)

	//通过指针对变量进行赋值
	*strPtr = "Java is Good too"
	fmt.Printf("str type is %T, value is %v, address is %v\n", str, str, &str)

	/*
	通过new函数创建了一个*string指针，并通过指针对其进行赋值。
	*/
	str1 := new(string)
	fmt.Printf("str1 type is %T. str1 value is %v，*str1 value is %v\n", str1, str1, *str1)
	*str1 = "Golang is Good"
	fmt.Printf("str1 type is %T. str1 value is %v，*str1 value is %v\n", str1, str1, *str1)

	//定义一个类型为*string，名称为surname的命令行参数
	//参数依次是命令行参数的名称，默认值，提示
	surname := flag.String("surname", "王", "您的姓")
	//定义一个类型为string，名称为personalName的命令行参数
	//除了返回指针类型结果，还可以直接传入变量地址获取参数值
	var personalName string
	flag.StringVar(&personalName, "personalName", "小二", "您的名")
	//定义一个类型为*int，名称为id的命令行参数
	id := flag.Int("id", 0, "您的ID")

	//解析命令行参数
	flag.Parse()

	fmt.Printf("I am %v %v, and my id is %v\n", *surname, personalName, *id)
}