package main

import "fmt"

func array() {
	//数组在函数传参中是传递的整个拷贝。
	//数组定义，
	// 1.默认为0填充，可以定义固定索引元素值，
	//数组是值语义，所以和c不同，不存在a + 1的使用方式。
	var a [3]int; 
	var b = [...]int {1, 2, 3}
	var c = [...]int {1, 3:4, 5}
	var d = [...]int {1, 2:3, 10:11}
	
	//数组指针除了类型和数组不同外，操作方式一样。都是采用索引的方式操作
	//函数传参可以采用数组指针方式，减少拷贝开销，并可以修改数组值。
	var pa =  & a; 
	fmt.Println(a[0], a[1])
	fmt.Println(b[0], b[1])
	for i, v: = range pa {
		fmt.Println(i, v)
	}
		
}