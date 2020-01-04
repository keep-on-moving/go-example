package main

import "fmt"

type rect struct {
	width,height int
}

//一般的函数定义叫做函数，定义在结构体上面的函数叫该结构体方法
func(r *rect)area()int{
	return r.width*r.height
}

func (r rect)perim() int {
	return 2*r.width+2*r.height
}
func main(){
	r := rect{width:10,height:5}
	fmt.Println("area:",r.area())
	fmt.Println("perim",r.perim())

	//Go会自动识别方法调用的参数是结构体变量还是结构体指针，如果你要修改结构体内部成员值，那么使用
	// 结构体指针作为函数限定类型，也就是说参数若是结构体 变量，仅仅会发生值拷贝。
	rp := &r
	fmt.Println(rp)
	fmt.Println("area:",rp.area())
	fmt.Println("perim",rp.perim())

	//从某种意义上说，方法是函数的“语法糖”。当函数与某个特定的类型绑定，那么它就是一个方法。也证因为如此，我们可以将方法“还原”成函数。
	//instance.method(args)->(type).func(instance,args)
	p :=Person{2, "张三"}
	p.test(1)
	var f1 func(Person, int) = Person.test
	f1(p,2)
	Person.test(p, 3)
	var f2 func(Person, int) = Person.test
	f2(p,4)

	pro:=Student{Person{2,"张三"},27}
	pro.retest()
}

type Person struct {
	Id int
	Name string
}

func(this Person)test(x int){
	fmt.Println("Id:",this.Id,"Name",this.Name)
	fmt.Println("x=",x)
}

type Student struct {
	Person
	Score int
}

func(this Person) retest(){
	fmt.Println("Preson retest")
}
//使用匿名字段，实现模拟继承。即可直接访问匿名字段（匿名类型或匿名指针类型）的方法这种行为类似“继承”。访问匿名字段方法时，有隐藏规则，这样我们可以实现override效果。
func (this Student) retest(){
	fmt.Println("student retest")
}