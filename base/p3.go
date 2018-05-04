// practice for type.

package main

import "fmt"
import "strconv"

func main() {
	var a int = 65
	b := string(a)
	fmt.Println(b)

	// sometime, i what print string `65`, not A, then how to?
	s := strconv.Itoa(a)
	c, _ := strconv.Atoi(s)
	fmt.Println("int => string, use: strconv.Itoa(), need import `strconv` for usage.")
	fmt.Println("ex: int[", a, "] => string[", s, "]")
	fmt.Println("string => int, use: strconv.Atoi(), need import `strconv` for usage.")
	fmt.Println("ex: string[", s, "] => int[", c, "]")
}
