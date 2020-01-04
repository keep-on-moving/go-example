package main

import "fmt"

//Go语言结构体数据类是将各个类型的变量定义的集合，通常用来表示记录。
type person struct {
	name string
	age int
}

func main(){
	fmt.Println(person{"bob",20})
	fmt.Println(person{name:"Alice",age:21})
	//未显示赋值的成员初始值为0
	fmt.Println(person{age:22})
	//&获取结构体变量的地址
	fmt.Println(&person{"Ann",50})
	s := person{"Sean",70}
	fmt.Println(s.name)

	//结构体指针也可以使用点号来访问结构体成员
	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

}