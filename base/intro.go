// intro for go.

package main

/* 1. for alias of package, usage: package.function */
import std "fmt"

/* 2. for simple of package, no need <package.><function>'s <package.> */
// 3. WARNING: `.` and `alias` MUST select one from two.
//import . "fmt"
import (
	"fmt"
	nb "fmt"
	// 4. you must use nb.<func> in the following, else will compile error.
)

// 5. const variable.
// const var.
const PI = 3.14
const (
	A = 1
	B = 2
	C = 3
)


// 6. global variable.
// global var.
var name = "gopher"
var (
	a = 1
	b = 2
	c = 3
)

// 7. declare type alias.
type myType int

// 8. declare struct hello.
type hello struct{}

// 9. declare interface world.
type world interface{}

type (
	newType int
	hi hello
	wd world
)

// 10. public func define: function name's 1st charactor MUST be uppercase.
func Public_func() {
	fmt.Printf("i am a public func, my function name's first charactor MUST be uppercase\n")
}

// 11. private func define: function name's 1st charactor MUST be lowercase.
func private_func() {
	fmt.Printf("i am a private func, my function name's first charactor MUST be lowercase\n")
}

//WARNING: main() must one, to compile: need 
//	a. package main
//	b. func main().
func main() {
	fmt.Println("go programming introduce")
	std.Printf("hello, i'm a alias for import 'fmt' to 'std'\n")
	nb.Printf("hello, import package name fmt is alias to `nb`\n")
	Public_func()
	private_func()
}

