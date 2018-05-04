// no support implicit cast

package main

func test() {
    var b byte = 100
    //> will be error, no support implicit cast.
    // var n int = b
    var n int = int(b)

    // use () for operation level.

    // no bool type not be a bool expression
    // the following will be compile error
    // a := 100
    // if a {
    //  ...
    //

    a := 100
    if a {
        println("true")
    }
}

func main() {
    test()
}
