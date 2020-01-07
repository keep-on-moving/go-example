package main

import (
	"encoding/base64"
	"fmt"
)

func main(){
	data := "abc23423423"
	//base64标准编码需要一个[]byte参数
	sEnc := base64.StdEncoding.EncodeToString([]byte (data))
	fmt.Println(sEnc)

	sDec,_ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	//兼容URL
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)

	uDec,_:= base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
