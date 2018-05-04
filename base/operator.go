//operator

package main

import "fmt"

func operator() {
	// operator: ^, $^...
	/* two operator nubmer.

	 6: 0110
	11: 1011
	--------
	 ^	1101 // the different is 1, the same is 0
	 &^ 0100 // clear bits, clear the second operator's number's bit.
	*/

	// ^: XOR是不进位加法计算，也就是异或计算
	// &^: 清空运算, 如果ybit位上的数是0则取x上对应位置的值， 如果ybit位上为1则取结果位上取0
	a := 6
	b := 11
	fmt.Println("a=", a, "b=", b, "a^b=", a^b)
	fmt.Println("a=", a, "^a=", ^a)
	fmt.Println("a=", a, "b=", b, "a&^b=", a&^b)

	var c uint32 = 6
	fmt.Println("c<uint32>=", c, "^c=", ^c)

	var d int8 = 1
	// ^d = -2.
	fmt.Println("d<int8>=", d, "^d=", ^d)

	var e uint8 = 1
	// ^e = 254.
	fmt.Println("e<uint8>=", e, "^e=", ^e)
}

func main() {
	fmt.Println("operator study...")
	operator()
}
