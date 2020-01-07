package main

import (
	"fmt"
	"os"
)

func ceateFile(p string) * os.File{
	fmt.Println("creating")
	f,err := os.Create(p)
	if err != nil{
		panic(err)
	}

	return  f
}

func writeFile(f *os.File){
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f * os.File){
	fmt.Println("close")
	f.Close()
}

func main(){

	f := ceateFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}
//Defer 用来保证一个函数调用会在程序执行的最后被调用，用来资源清理工作

