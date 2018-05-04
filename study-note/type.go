package main

import "unsafe"

var x int
var f float32 = 1.6
var s = "abc"
var a = 1
var p, q, r int
var d, e = "abc", 111


var (
    g int
    h float32
)

func hello() {
    //1. local variable cover global variable.
    x := 123
    var b = 13.4
    var a = "hello, world, i am a local variable, there is a global variable a is integer."
    print(x)
    print(b)
    print(a)
    print(d)
    println("")
    println(x, b)
    println(a)
    println(d)

    //2. multi-declare or assign, use old value, and calc.
    data, i := [3]int{1, 2, 3}, 0
    i, data[i] = 2, 100 // i = 0, data[0] = 100, i= 2
    println(i, data[0], data[1], data[2])
}

func test() (int, string) {
    return 1, "abc"
}

func get() {
    //3. '_' is for ignore usage.
    _, s = test()
    println(s)
    i := 10
    // you can use _ to avoid complie warnning.
    // you can comment the following or uncomment it. and compile file.
    _ = i
}

func re_assign() {
    s := "abc"
    println(&s, s)

    //4. s is re-assign and y is declare.
    // string re-assign the address is equal.
    s, y := "hello", 20
    println(&s, y)
    {
        s, z := 1000, 30
        println(&s, z)
    }
}

// const variable.
func const_var() {
    const x, y int = 1, 2
    const s = "hello, world!"

    const (
        a, b = 10, 100
        c bool = false
    )

    const (
        d = "hello, fucking."
        e // default is d
    )
    println(x, y, a, b, c)
    println(s)

    // you can't get the const variable address.
    // the two following line will compile error.
    //println(&d, d)
    //println(&e, e)
    println(d)
    println(e)


    // const variable value can be value when compiling is sure.
    const (
        f = "world"
        g = len(f)
        // unsafe is compile error, what the fucking in go 1.9???
        //done: import unsafe...
        h = unsafe.Sizeof(g)
    )
    println(f, g, h)

    // const variable will never overflows.
    const (
        i byte = 100
        //j int = 1e20 // float64 to int, overflows
    )

    println(a, b)
}

func week() {
    const (
        sunday = iota  // will be 0, the following will add one automally.
        monday
        tuesday
        wednesday
        friday
        saturday
    )

    print(sunday, monday, tuesday, wednesday, friday, saturday)
}

func bytes() {
    const (
        _ = iota // iota = 0
        kb int64 = 1 << (10 * iota) // iota = 1
        mb       // iota = 2.
        gb
        tb
    )
    println("")
    println(kb)
    println(mb)
    println(gb)
    println(tb)
}

func iota_test() {
    const (
        a, b = iota, iota << 10
        c, d
    )
    println(a)
    println(b)
    println(c)
    println(d)

    const (
        A = iota
        B
        C = "C"
        D
        E = iota
        F
    )
    println(A)
    println(B)
    println(C)
    println(D)
    println(E)
    println(F)
}

func main() {
    hello()
    get()
    re_assign()
    const_var()
    week()
    bytes()
    iota_test()
}
