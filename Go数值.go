package main

import "fmt"

func main(){
	//字符串可以使用+号连接
	fmt.Println("go"+"lang")
	fmt.Println("1+1=",1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)

	//布尔型的几种操作符
	fmt.Println(true && false)

	fmt.Println(true || false)

	fmt.Println(!true)
}