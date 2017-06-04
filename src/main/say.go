package main

import (
	"fmt"
)

var( //全局变量
	goo string = "daaa"
	doo int = 233
)

const A string  = "abd" //常量字符串
const B   = "阿打算打算" //常量字符串 隐式

func main(){
    	//局部变量
	var name  = "danxingxi"
	var name1  = "go"
	var name2  = "233 go"
	var aa, ab, ac  = "adad","adcsw", "asdasd"

	fmt.Println( name )

	fmt.Println("hello golang " + name , name1, name2)

	fmt.Println(aa , ab, ac)

	fmt.Println(goo, doo)

	var numb  = 22
	var point *int /* 声明指针变量 */

	point = &numb  /* 指针变量的存储地址 */

	fmt.Println("变量地址：", &numb)

	fmt.Println("point 变量储存的指针地址:：", point)

	fmt.Println("使用指针访问值:", *point)

	fmt.Println("point不是空指针：",point!=nil)

}