// reference type.

package main

import "unsafe"

func test() {
    a := []int{0, 1, 2}
    a[1] = 10
    println(a, unsafe.Sizeof(a))

    b := make([]int, 3)
    b[1] = 10
    println(b, unsafe.Sizeof(b))

    c := new([]int)
    //the following will be compile error..
    // return is a type *[]int...
    //c[1] = 10
    println(c, unsafe.Sizeof(c))
}

func main() {
    test()
}
