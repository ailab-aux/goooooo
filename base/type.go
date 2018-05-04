// for base type.
package main

import "fmt"
import "math"

// base type.
// bool true/false
// int8/uint8 int16/uint16 int32/uint32 int64/uint64
// int: 32/64 for os.
// float32/float64
// byte, alias of uint8
// rune, alias of int32
var a bool = false

var b int8 = 127

//
var c int = 333

// system will recongnize it type.
var d = 444

var e byte = 255
var f rune = 1024

var g1, g2, g3, g4, g5 int = 1, 2, 3, 4, 5
var h1, h2, h3, h4, h5 = 1, 2, 3, 4, 5

// complex type: complex64/complex128
var x complex64

var y int64 = math.MaxInt64

// 64bit system, int is a int64, 32bit system int is a int32.
var z int = math.MaxInt64

func assign() {
	var x int
	x = 12
	var y int = 13
	var z = 14
	z1 := 15

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(z1)

	var a1, b1, c1, d1, e1 int = 1, 2, 3, 4, 5
	var j1, j2, j3, j4, j5 = 1, 2, 3, 4, 5
	l1, l2, l3, l4, l5 := 1, 'a', "string", 1.1, 3000

	fmt.Printf("a1=%d,b1=%d,c1=%d,d1=%d,e1=%d\n", a1, b1, c1, d1, e1)
	fmt.Printf("j1=%d,j2=%d,j3=%d,j4=%d,j5=%d\n", j1, j2, j3, j4, j5)
	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(l3)
	fmt.Println(l4)
	fmt.Println(l5)

	// ff is a float64 type.
	var ff = 1024.1024
	fmt.Println(ff)
}

func cast() {

	fmt.Println("float64 <=> int")
	// must implicit cast, no explicit cast.
	var f1 = 1.1
	var i1 = int(f1)
	fmt.Print(f1, i1)
	var fx float64 = float64(i1)
	fmt.Print("\n", fx, i1)

	fmt.Println("\ncustom type[type newType int] <=> int, need implicit cast")
	//custom type.
	type newType int
	// explicit cast: int => newType. complie error!!!
	//var i2 int = 33
	//var n1 newType = i2 // error or not?
	//fmt.Print(i2, n1)

	fmt.Println("\ncustom type[type newType int] <=> int, need implicit cast")
	var n2 newType = 44
	// implicit cast: int => newType, complie ok.
	var i3 int = int(n2)
	fmt.Print(n2, i3)
	i3 = int(n2)
	fmt.Print("\n", n2, i3)

	fmt.Println("\nint => string")
	// int => string
	var i4 = 66
	var s1 = string(i4)
	fmt.Print(i4, s1)

	fmt.Println("\nstring => int can't use implicit cast, need what?")
	//var s2 = "B"
	//var ii = int(s2)
	//fmt.Print("\n", ii, s2)

	fmt.Println("\nbyte <=> uint8, no need implicit cast, is the same type.")
	// byte => uint8, no need implicit cast.
	var b1 byte = 127
	var u1 uint8 = b1
	var b2 byte = u1
	fmt.Print(b1, u1, b2)

	fmt.Println("\nrune <=> int32, no need implicit cast, is the same type.")
	// int32 => rune, no need implicit cast.
	var r1 rune = 1024
	var i5 int32 = r1
	var r2 rune = i5
	fmt.Print(r1, i5, r2)
}

func main() {
	fmt.Println(a)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println("call assign() function")
	assign()
	cast()
}
