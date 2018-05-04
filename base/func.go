// function.

package main

import "fmt"

func f1(x int) int {
	return x
}

func f2(x ...int) int {
	return 2
}

func f3(s string, x ...int) int {
	return 3
}

func f4(x, y, z int) (int, int, int) {
	return x, y, z
}

func f5(x, y, z int) (a, b, c int) {
	a, b, c = x, y, z
	return
}

func f6(x, y, z int) (a, b, c int) {
	a, b, c = x, y, z
	return a, b, c
}

func f7(x []int) {
	x[0] = 5
	x[1] = 6
}

func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func f8() {
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}

func df() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")

	// i is a value copy.
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

	// i is a reference.
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func panic_and_recover() {
	fmt.Println("will painc at defer by func, the func will recover the painc.")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("hello, i am recover the panic...")
		}
	}()

	panic("i am panic...")
	fmt.Println("i am never execute...")
}

func main() {
	fmt.Println("call f1(x int) int ...")
	fmt.Println(f1(10))

	s := []int{1, 2, 3}
	fmt.Println("call f2(x...int) int ...")
	fmt.Println(f2(s[1], s[2]))

	fmt.Println("call f3(s string, x...int) int ...")
	fmt.Println(f3("a", 1))

	fmt.Println("call f4(x, y, z int) (int, int, int)...")
	fmt.Println(f4(1, 2, 3))

	//a, b, c := f5(2, 3, 4)
	fmt.Println("call f5(x, y, z int) (a, b, c int)...")
	fmt.Println(f5(2, 3, 4))

	fmt.Println("call f6(x, y, z int) (a, b, c int)...")
	fmt.Println(f6(2, 3, 4))

	fmt.Println("call f7(x []int) ...")
	x := []int{1, 2}
	fmt.Println(x)
	f7(x)
	fmt.Println(x)

	// array is not a slice. will error.
	//y := [...]int{1, 2, 3}
	//fmt.Println(y)
	//f7(y)
	//fmt.Println(y)

	fmt.Println("closure ...")
	f8()

	fmt.Println("defer")
	df()

	panic_and_recover()
}
