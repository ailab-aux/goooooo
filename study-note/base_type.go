// base type.

package main

import "unsafe"
import "math"

func test() {
    var b bool
    println(b, unsafe.Sizeof(b))

    var bt byte
    println(bt, unsafe.Sizeof(bt))


    var r rune
    println(r, unsafe.Sizeof(r))

    var i int
    println(i, unsafe.Sizeof(i))

    var ui uint
    println(ui, unsafe.Sizeof(ui))

    var i8 int8
    println(i8, unsafe.Sizeof(i8))

    var u8 uint8
    println(u8, unsafe.Sizeof(u8))

    var i16 int16
    println(i16, unsafe.Sizeof(i16))

    var u16 uint16
    println(u16, unsafe.Sizeof(u16))

    var i32 int32
    println(i32, unsafe.Sizeof(i32))

    var u32 uint32
    println(u32, unsafe.Sizeof(u32))

    var i64 int64
    println(i64, unsafe.Sizeof(i64))

    var u64 uint64
    println(u64, unsafe.Sizeof(u64))

    var f32 float32
    println(f32, unsafe.Sizeof(f32))

    var f64 float64
    println(f64, unsafe.Sizeof(f64))

    //var c complex
    //println(c, unsafe.Sizeof(c))

    var c64 complex64
    println(c64, unsafe.Sizeof(c64))

    var c128 complex128
    println(c128, unsafe.Sizeof(c128))

    var up uintptr
    println(up, unsafe.Sizeof(up))

    //var a array
    //var s struct
    //var s string
    //var s slice
    //var m map
    //var c channel
    //var i interface
    //var f function

    v1, v2, v3, v4 := 071, 0x1f, 1e9, math.MinInt16
    println(v1)
    println(v2)
    println(v3)
    println(v4)

    // the following will be error, wtf???
    //var p uintptr = nil
    var p uintptr
    println(p, unsafe.Sizeof(p))
}

func main() {
    test()
}
