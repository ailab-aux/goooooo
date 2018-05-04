// array.

package main

import "fmt"

func array() {
	var a [2]int //default is {0, 0}
	var b [1]int
	// a's type is NOT equal b's type.
	fmt.Println(a)
	fmt.Println(b)

	c := [2]int{}
	d := [1]int{}
	fmt.Println(c)
	fmt.Println(d)

	// auto get length.
	e := [...]int{1, 2, 3, 4, 5}
	f := [...]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	g := [...]int{99: 1}
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// array's compare, only support: ==, !=
	// two's type must be the SAME.
	cmp_a := [3]int{1, 2, 3}
	cmp_b := [3]int{1, 2, 3}
	cmp_c := [3]int{0, 2, 3}
	if cmp_a == cmp_b {
		fmt.Println("cmpare a=", cmp_a, "==", "b=", cmp_b)
	}
	if cmp_a != cmp_c {
		fmt.Println("cmpare a=", cmp_a, "!=", "c=", cmp_c)
	}

	// array's pointer
	pa := &a
	var pb *[1]int = &b
	fmt.Println("a=", a, "pa=", pa)
	fmt.Println("b=", b, "pb=", pb)

	// pointer array.
	x, y := 1, 2
	var ap = [2]*int{&x, &y}
	fmt.Println(ap)

	// new return a pointer of array.
	na := new([3]int)
	na[1] = 1
	fmt.Println(na)

	// multi array.
	ma := [3][2]int{
		{0, 0},
		{1, 1},
		{2, 2},
	}
	fmt.Println("multi array ma=[3][2]int=", ma)

	mb := [...][3]int{
		{0, 1, 2},
		{1, 2, 3},
	}
	fmt.Println("multi array mb=[...][3]int=", mb)
}

func main() {
	array()
}
