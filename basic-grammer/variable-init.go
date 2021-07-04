package main

import "fmt"

func getName() (string, string) {
	return "王", "小二"
}

//变量初始化
func main() {
	var a int = 100
	var b = "100"
	c := 0.17

	fmt.Printf("a value is %v, type is %T\n", a, a)
	fmt.Printf("b value is %v, type is %T\n", b, b)
	fmt.Printf("c value is %v, type is %T\n", c, c)

	var d = 1
	var e = 2
	d, e = e, d
	fmt.Printf("d = %v, e = %v\n", d, e)

	surName, _ := getName()
	_, personalName := getName()
	fmt.Printf("My surname is %v and my personal name is %v\n", surName, personalName)
}

