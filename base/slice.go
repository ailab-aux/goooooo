// slice
package main

import "fmt"

func slice() {
	var s1 []int
	fmt.Println(s1)

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)

	// slice one element.
	s2 := a[9]
	fmt.Println("slice one element")
	fmt.Println(s2)

	// slice [begin, end)
	s3 := a[5:10]
	fmt.Println("slice multi element[begin:end], range[begin,end)")
	fmt.Println(s3)

	// slice range: [begin:len(a)]
	s4 := a[5:len(a)]
	fmt.Println("slice multi element[begin:len(a)], range[begin, end)")
	fmt.Println(s4)

	// slice range: [begin:]
	s5 := a[5:]
	fmt.Println("slice multi element[begin:], range[begin, end)")
	fmt.Println(s5)

	// slice range: [:end]
	s6 := a[:5]
	fmt.Println("slice multi element[:end], range[0, end)")
	fmt.Println(s6)

	// make a slice.
	// make(type, len, cap)
	s7 := make([]int, 3, 10)
	fmt.Println("slice create by make(type, len, cap)")
	fmt.Println(s7, "len=", len(s7), "cap=", cap(s7))

	// make(type, len, cap)
	s8 := make([]int, 4)
	fmt.Println("slice create by make(type, len), cap will be equal as len.")
	fmt.Println(s8, "len=", len(s8), "cap=", cap(s8))

	//NOTE: when you memory is not enough, will incre cap to 2*cap caps.

	b := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	sb := b[2:5]
	fmt.Println("sb=b[2:5] =", string(sb), "len =", len(sb), "cap =", cap(sb))

	// re-slice. from slice to slice.
	sc := sb[1:3]
	fmt.Println("sc=sb[1:3] =", string(sb), "len =", len(sc), "cap =", cap(sc))

	//WARNING: if you index is over range, will error.

	// append.

	fmt.Println("slice append...")

	s9 := make([]int, 3, 6)
	fmt.Printf("%p, %d\n", s9, s9)
	s9 = append(s9, 1, 2, 3)
	fmt.Printf("%p, %d\n", s9, s9)
	s10 := append(s9, 4, 5, 6)
	fmt.Printf("%p, %d, address changed, larger than cap, will re-alloc cap=%d\n", s10, s10, cap(s10))
	fmt.Println("when address changed, origin slice and new slice will different address space, change content will be no effective.")
	s9[0] = 9
	s10[0] = 10
	fmt.Println("origin slice change[0]=>9, all content =", s9)
	fmt.Println("new slice change[0]=>10, all content =", s10)

	// copy.
	fmt.Println("array copy(dst, src)...")
	sx := []int{1, 2, 3, 4, 5, 6}
	sy := []int{7, 8, 9}
	fmt.Println("copy(x =", sx, ", y =", sy, ")")
	copy(sx, sy)
	fmt.Println("dest x =", sx)
	sm := []int{1, 2, 3, 4, 5, 6}
	sn := []int{7, 8, 9}
	fmt.Println("copy(n =", sn, ", m =", sm, ")")
	copy(sn, sm)
	fmt.Println("dest n =", sn)

	fmt.Println("copy will as dest array/slice length.")

	// slice copy, by range.
	//todo::
	sp := []int{1, 2, 3, 4, 5, 6}
	sq := []int{9, 8, 7, 6, 5, 4}
	fmt.Println("copy(p[1:3] =", sp[1:3], ", q[4:6] =", sq[4:6], ")")
	copy(sp[1:3], sq[4:6])
	fmt.Println("p =", sp)

	// practice.
	z := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	sz := z[:]
	fmt.Println("z =", z)
	fmt.Println("sz:=z[:] =", sz)
}

func main() {
	slice()
}
