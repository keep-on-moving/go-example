package main

import (
	"fmt"
)

func main(){
	i := 1
	for i<=3 {
		fmt.Println(i)
		i++
	}

	for j:=7;j<=9;j++{
		fmt.Println(j)
	}

	for{
		fmt.Println("loop")
		break
	}

	a := strrev("abcd");
	fmt.Println(a)

	b:=strrev1("abcd")
	fmt.Println(b)
}

//直接头尾互换
func strrev(s string) string{
	byteArr := []byte(s)
	str := ""
	for i:= len(s)-1;i>=0;i--{
		str = str + string(byteArr[i])
	}

	return  str
}

//中间区中 然后左右互换
func strrev1(s string) string{
	byteArr := []byte(s)
	for i:=0;i<len(byteArr)/2;i++{
		byteArr[i],byteArr[len(byteArr)-1-i] = byteArr[len(byteArr)-1-i], byteArr[i]
	}
	return string(byteArr)
}
