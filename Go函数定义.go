package main

import "fmt"

func main(){
	res := plus(100,222)

	fmt.Println(res)
}

//函数声明 两个int参数 返回值亦为int
func plus(a int, b int)int{
	return a+b
}
