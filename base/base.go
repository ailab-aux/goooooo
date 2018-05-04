// statement base.

package main

import "fmt"

func pointer() {
	var a = 15
	var pa = &a
	fmt.Println("a=", a, "pa=", pa, "*pa=", *pa)

	var b int = 1024
	var pb *int = &b
	fmt.Println("b<int>=", b, "pb=", pb, "*pb<int>=", *pb)

	c := "hello, world"
	pc := &c
	fmt.Println("c=", c, "pc=", pc, "*pc=", *pc)
}

func add_one() {
	a := 1
	// a++ must one statement, NOT a expression, can't use like: b := a++.
	// no ++a
	a++
}

func larger_0(x int) bool {
	r := x > 0
	return r
}

func if_state() {
	// if usage 1.
	a := 1
	if a >= 1 {
		fmt.Println("if a >= 1 matched...")
	}

	// if usage 2.
	if b := 0; b >= 0 {
		fmt.Println("if b := 0; b >= 0 matched, you can use a initialize expression: b := 0; ...")
		fmt.Println("NOTE: initialize variable b's scope is in if statement...")
	}

	// if usage 3.
	if c := larger_0(-1); c {
		fmt.Println("call x(-1)'s return larger than 0...")
	} else {
		fmt.Println("call x(-1)'s return is not larger than 0...")
	}
}

func for_state() {
	// for usage 1.
	a := 1
	fmt.Println("for usage: loop forever.")
	for {
		a++
		if a > 10 {
			break
		}
		fmt.Print(a, " ")
	}
	fmt.Print("\n")

	// for usage 2.
	fmt.Println("for usage: one condition state")
	for a > 1 {
		a--
		fmt.Print(a, " ")
	}
	fmt.Print("\n")

	// for usage 3.
	fmt.Println("for usage: initialize; condition expression; break condition")
	fmt.Println("NOTE: initialize expression variable's scope is in for statement...")
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Print("\n")

	// NOTE: condition expression if need calc, you should calc before for statement.
	s := "hello, world"
	l := len(s)
	for i := 0; i < l; i++ {
		fmt.Print(string(s[i]))
	}
	fmt.Println("\ncondition expression SHOULD calc before for statement.")
}

func switch_state() {
	// switch usage 1:
	fmt.Println("switch usage: case const variable...")
	fmt.Println("NOTE: case no break...")
	fmt.Println("NOTE: case goto next, use keyword: fallthrough ")
	a := 10
	switch a {
	case 1:
		fmt.Println("a=1")
	case 2:
		fmt.Println("a=2")
	case 3:
		fmt.Println("a=3")
	case 4:
		fmt.Println("a=4")
	case 5:
		fmt.Println("a=5")
	default:
		fmt.Println("default: a=", a)
	}

	// switch usage 2:
	fmt.Println("\nswitch usage: case is a expression...")
	switch {
	case a > 0:
		fmt.Println("a > 0")
		fallthrough
	case a > 1:
		fmt.Println("a > 1")
		fallthrough
	case a >= 10:
		fmt.Println("a >= 10")
	default:
		fmt.Println("this is WHAT condition will match? a=", a)
	}

	// switch usage 3:
	fmt.Println("\nswitch usage: switch initialize, case is const or expression...")
	fmt.Println("the following case is const variable")
	x := 10
	switch b := larger_0(x); {
	case true:
		fmt.Println("x is larger 0, b=", b)
	case false:
		fmt.Println("x is less or equal 0, b=", b)
	}
	// don't allow define a function in function.
	//func larger_zero(x int) bool {
	//  return x > 0
	//}

	fmt.Println("the following case is expression")
	x = 0
	switch b := larger_0(x); {
	case b == true:
		fmt.Println("x is larger 0")
	case b == false:
		fmt.Println("x is less or equal 0")
	}
}

func break_state() {
	// break usage.
	fmt.Println("break state usage...")
OUTSIDE:
	for {
	INSIDE:
		for {
			fmt.Println("break inside")
			break INSIDE
		}
		fmt.Println("break outside")
		break OUTSIDE
	}
	fmt.Println("break 1 OK")

	for {
		fmt.Println("break one level loop...")
		break
	}
	fmt.Println("break 2 OK")
}

func continue_state() {
	// continue usage.
	fmt.Println("continue state usage...")
OUTSIDE:
	for i := 0; i < 3; i++ {
		fmt.Println("continue outside loop, i=", i)
	INSIDE:
		for j := 0; j < 2; j++ {
			if j == 1 {
				continue INSIDE
			}
			fmt.Println("continue inside loop, j=", j)
		}
		if i == 2 {
			continue OUTSIDE
		}
	}
	fmt.Println("continue case 1 over...\n")
	for i := 0; i < 3; i++ {
		if i == 2 {
			continue
		}
		fmt.Println("continue, one level for loop, i=", i)
	}

	fmt.Println("exit continue_state()")
}

func goto_state() {
	// goto usage.
	fmt.Println("goto statement usage...")
	for i := 0; i < 10; i++ {
		if i == 9 {
			goto OUT
		}
		fmt.Println("for loop inside, i=", i)
	}
	fmt.Println("before goto OUT")
OUT:
	fmt.Println("goto OUT")
}

func main() {
	pointer()
	if_state()
	for_state()
	switch_state()
	break_state()
	continue_state()
	goto_state()
}
