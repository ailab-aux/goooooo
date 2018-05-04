// for storage unit enum.

package main

import "fmt"

const (
	// if type is uint64, will overflow...
	Byte float64 = 1 << (10 * iota)
	KByte
	MByte
	GByte
	TByte
	PByte
	EByte
	ZByte
	YByte
)

func main() {
	fmt.Println("byte=", Byte)
	fmt.Println("k-byte=", KByte)
	fmt.Println("m-byte=", MByte)
	fmt.Println("g-byte=", GByte)
	fmt.Println("t-byte=", TByte)
	fmt.Println("p-byte=", PByte)
	fmt.Println("e-byte=", EByte)
	fmt.Println("z-byte=", ZByte)
	fmt.Println("y-byte=", YByte)
}
