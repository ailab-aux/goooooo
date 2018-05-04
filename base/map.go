// map stucy.

package main

import "fmt"

func map_study() {
	var m map[int]string
	m = map[int]string{}
	fmt.Println(m)
	m[1] = "ok"
	fmt.Println(m)

	var m1 map[int]string = make(map[int]string)
	m2 := make(map[int]string)
	fmt.Println(m1, m2)

	m1[5] = "hello"
	fmt.Println(m1)

	m2[1024] = "world"
	fmt.Println(m2)

}

func main() {
	map_study()
}
