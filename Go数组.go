package main

import (
	"fmt"
)

func main(){
	//定义五个元素的整型数组
	var a [5]int
	fmt.Println("emp:", a)

	a[4] =100
	fmt.Println("set:", a)
	fmt.Println("get:",a[4])

	//内置len函数计算数组的长度
	fmt.Println("len:",len(a))

	b := "ssfefefeff"
	fmt.Println("len:",len(b))

	c:=[5]int{1,2,3,4,5}
	fmt.Println("dcl:", c)

	//数组都是一维的，但是我们可以把数组的元素定义为一个数组
	//来获取多维数组结构
	var twoD [2][3][4]int
	for i :=0;i<2 ; i++ {
		for j:=0;j<3;j++{
			for k:=0;k<4;k++{
				twoD[i][j][k] = i+j+k
			}
		}
	}

	fmt.Println("2d:", twoD)

	//new 创建数组，返回数组的指针
	var e = new([5]int)

	fmt.Println(e,len(e))

	aa := [...]User{
		{0, "User0"},
		{8, "User8"},
	}
	fmt.Println(aa)

	//返回User的 物理地址
	bb := [...]*User{
		{0, "User0"},
		{8, "User8"},
	}

	fmt.Println(bb)
}

type User struct {
	Id int
	Name string
}


