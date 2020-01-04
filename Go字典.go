package main

import "fmt"

func main(){
	//创建一个字典可以使用内置函数make
	//make(map[键类型]值类型)
	m := make(map[string]int)
	m["K1"] = 7
	m["K2"] = 13

	fmt.Println("map:", m)

	v1 := m["K1"]
	fmt.Println(v1)

	fmt.Println("len:",len(m))

	delete(m,"K2")
	fmt.Println("map:", m)

	//根据键来获取值有一个可选的返回值，第一个值为该键位的value，第二个就是有没有这个
	a,b :=m["K1"]
	fmt.Println(a,b)

	_,c := m["k2"]
	fmt.Println(c)
	//:= 初始化字典
	n := map[string]int{"foo":1,"bar":2}
	fmt.Println("map:",n)

}
