// const expression.
package main

import "fmt"

const MAX_COUNT int = 12
const MAX = 13

var x = 13

const (
	a = 'a'
	b = 2
	c = MAX
	d = MAX_COUNT
	//complie error, const can't assign by global var.
	//e = x
)

const (
	//first charactor upper for public, lower for private.
	A_public  = 10
	b_private = 15
)

// const attribute. only const have? The answer is YES.
const (
	attr_a = 1
	// will be 1
	attr_b
	// will be 1
	attr_c

	attr_d, attr_e = 2, 3
	// following must be `<const var>, <const var>` else compile error.
	attr_f, attr_g

	// can be the different type.
	attr_h, attr_i = "hello", 1024.1024
	attr_j, attr_k
)

// iota for count at begin.
const (
	// count_a = 0
	count_a = iota
	count_b = 'b'
	// count_c = 2
	count_c = iota
	// count_d = 3
	count_d
	// count_d = 4
	count_e
)

var (
	//global variable can't use iota for counter.
	//g_a = iota
	g_b = 'b'
	//g_c = iota
	g_d = 0

	G_public  = "hello, world"
	g_private = "hi, world"

	g_e, g_f, g_g, g_h = 1, 2, '3', "hello"
	// the following will compile error, global var can't use the attribute of const.
	//g_i, g_k, g_l, g_m
)

const (
	con_a, con_b, con_c, con_d = 1, "b", "hello", 1.9
)

func main() {
	fmt.Println("hello, const assignment and `iota` usage, and continous assignment.")
}
