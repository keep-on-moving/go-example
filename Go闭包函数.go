package main

import "fmt"

func intSeq() func() int{
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func closure(x int) func (y int) int {
	return func(y int) int {
		return x + y
	}
}

func adder() func(int) int{
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

//闭包找到的是同一地址中父级函数中对应变量最终的值
//https://blog.csdn.net/weixin_43586120/article/details/89456183  闭包解析
func main(){
	//返回为 i=0的方法体
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	//上面函数体一直在 所以i一直在添加 而当重新定义时，则相当于重建另外一个方法体
	newInt := intSeq()
	fmt.Println(newInt())

	//返回为Y为参数的函数 X为10
	add10 := closure(10)
	fmt.Println(add10(5))
	fmt.Println(add10(6))

	add20 := closure(20)
	fmt.Println(add20(5))

	var fs []func() int
	for i:=0;i<3;i++{
		fs = append(fs,func() int{
			return  i
		})
		fmt.Println(fs)
	}
	for _,f:= range fs{
		fmt.Println("%p = %v \n", f, f())
	}

	result := adder()
	for i :=0 ; i<10; i++ {
		fmt.Println(result(i))
	}
}

