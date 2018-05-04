// for strong enum type.

package main

type Color int

const (
    Black Color = iota
    Red
    Blue
    White
    Yellow
    Green
)

func test(c Color) {}

func main() {
    c := Black
    test(c)

    // the following will compile error, can change int to Color.
    //x := 1
    //test(x)

    // compile ok, const will be change to Color...
    test(1)
}
