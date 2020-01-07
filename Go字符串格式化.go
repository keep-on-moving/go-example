package main

import (
	"fmt"
	"os"
)
type ponit struct {
	x,y int
}

func main (){
	p := ponit{1,2}
	//打印结构体对象的值 {1 2}
	fmt.Println("%v\n", p)

	//打印结构体的成员名称和值 {x:1 y:2}
	fmt.Println("%+v\n", p)

	//%#v 输出一个Go语法表示方式  main.point{x:1,y:2}
	fmt.Println("%#v\n", p)

	//%T 用来输出一个值的数据类型  main.point
	fmt.Println("%T\n", p)

	//%t格式化布尔型变量  true
	fmt.Println("%t\n", true)

	//%d 十进制格式化整型   123
	fmt.Println("%d\n", 123)

	//%b 二进制格式化  1110
	fmt.Println("%b\n", 14)

	//%c 整型对应的字符
	fmt.Println("%c\n", 33)

	//%x输出16位进制
	fmt.Println("%x\n", 456)

	//%s 输出基本的字符串
	fmt.Println("%s\n", "\"string\"")

	//%q 源码输出
	fmt.Println("%q\n","\"string\"")

	//%p输出指针的值
	fmt.Println("%p\n", &p)

	//%s 格式化后字符串赋值给一个变量
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}