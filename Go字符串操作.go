package main

import  s "strings"
import "fmt"

//给fmt.Println起别名方便调用
var p = fmt.Println

func main(){
	//是否包含字串
	p("Contains: ", s.Contains("test", "es"))
	//出现次数
	p("count: ", s.Count("test", "t"))
	//前缀
	p("HasPrefix", s.HasPrefix("test", "te"))
	//后缀
	p("hassuffix", s.HasSuffix("test", "st"))
	//判断首次出现位置
	p("index: ",s.Index("test", "s"))
	//join  是每个元素拼接
	p("join: ", s.Join([]string{"a","b"}, "_"))
	//repeat 重复
	p("repeat: ", s.Repeat("sss", 3))
	//replace  替换  -1替换所有
	p("Replace:", s.Replace("foo", "o", "sss", 1))
	p("Replace: ", s.Replace("fooo", "o", "dd", -1))
	p("Split:", s.Split("a-v-e-re-re", "-"))
	p("ToLower", s.ToLower("TEST"))
	p("ToUpper", s.ToUpper("test"))

	p("Len : ", len("hello"))
	//TODO  最后这个我不是很理解
	p("Char:", "hello"[2])
}
