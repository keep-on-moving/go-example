package main

import "sort"
import "fmt"

type ByLength [] string

func (s ByLength)Len() int{
	return  len(s)
}

func (s ByLength) Swap(i, j int){
	s[i],s[j] = s[j], s[i]
}

func (s ByLength) Less (i,j int) bool {
	return len(s[i]) < len(s[j])
}

//实现sort接口的 Len,Less和 Swap方法   Sort包的通用方法Sort
func main(){
	fruits := []string{"aaa", "bbbbbb", "ddddddddddddd"}
	a := ByLength(fruits)
	fmt.Println(a)
	sort.Sort(a)
	fmt.Println(fruits)
}