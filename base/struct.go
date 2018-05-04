// struct.

package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func assign() {
	a := person{}
	a.Age = 19
	a.Name = "james"

	fmt.Println(a)
	b := person{
		Age:  20,
		Name: "harden",
	}
	fmt.Println(b)
}

func by_value(p person) {
	p.Age = 200
	p.Name = "value"
	fmt.Println("call over, by_value(p person), p.Age =", p.Age, ", p.Name =", p.Name)
}

func by_ref(p *person) {
	p.Age = 2000
	p.Name = "reference"
	fmt.Println("call over, by_ref(p person), p.Age =", p.Age, ", p.Name =", p.Name)
}

func func_arg() {
	a := person{
		Age:  21,
		Name: "a",
	}
	fmt.Println("a is a value type...")
	fmt.Println("before call, by_value(p person), a.Age =", a.Age, ", a.Name =", a.Name)
	by_value(a)
	fmt.Println("after call, by_value(p person), a.Age =", a.Age, ", a.Name =", a.Name)

	fmt.Println("before call, by_ref(p person), a.Age =", a.Age, ", a.Name =", a.Name)
	by_ref(&a)
	fmt.Println("after call, by_ref(p person), a.Age =", a.Age, ", a.Name =", a.Name)

	b := &person{
		Name: "b",
		Age:  22,
	}
	fmt.Println("b is a ref type...")
	fmt.Println("before call, by_value(p person), b.Age =", b.Age, ", b.Name =", b.Name)
	by_value(*b)
	fmt.Println("after call, by_value(p person), b.Age =", b.Age, ", b.Name =", b.Name)

	fmt.Println("before call, by_ref(p person), b.Age =", b.Age, ", b.Name =", b.Name)
	by_ref(b)
	fmt.Println("after call, by_ref(p person), b.Age =", b.Age, ", b.Name =", b.Name)
}

type inner_anonymous struct {
	Age     int
	Name    string
	Contact struct {
		Phone string
		City  string
	}
}

type name_anonymous struct {
	int
	string
}

func anonymous_struct() {
	a := struct {
		Age  int
		Name string
	}{
		Name: "a",
		Age:  200,
	}
	fmt.Println("anonymous struct a by value: ", a)

	b := &struct {
		Age  int
		Name string
	}{
		Age:  2000,
		Name: "blablabla...",
	}
	fmt.Println("anonymous struct b by ref: ", b)

	c := inner_anonymous{
		Age:  1000,
		Name: "inner anonymous...",
	}

	c.Contact.Phone = "111111111111"
	c.Contact.City = "BeiJing"

	fmt.Println("inner anonymous assign, c: ", c)

	d := name_anonymous{100, "name anonymous..."}
	fmt.Println("struct's inner name anonymous, d: ", d)
}

type person_a struct {
	Age  int
	Name string
}
type person_b struct {
	Age  int
	Name string
}

func struct_assign_and_compare() {
	// same type assign.
	a := person_a{Age: 10, Name: "i am a struct person_a"}
	var b person_a
	b = a
	fmt.Println("a's type=person_a, b's type=person_a, can assign b=a, a=", a, ", b=", b)
	fmt.Println("a, b is the same type, can compare equal or not equal, a == b?", a == b, ", a != b?", a != b)
	var c = person_a{Age: 100, Name: "i am a person_a c"}
	fmt.Println("c =", c)
}

// compose not a inhert, go no supported inhert.
type human struct {
	Name string
	Age  int
	Sex  bool
}
type teacher struct {
	human
	course string
}
type student struct {
	human
	score int
}

type fucker struct {
	human
	Name string
}

func compose() {
	a := teacher{course: "math", human: human{Name: "james lebon", Age: 32, Sex: true}}
	b := student{score: 100, human: human{Name: "james harden", Age: 27, Sex: true}}

	fmt.Println("a is a compise struct, =", a)
	fmt.Println("b is a compise struct, =", b)

	a.human.Age = 100
	a.Age = 99
	a.human.Sex = false
	a.Sex = false

	fmt.Println("a go to korea to make a face change, a =", a)

	// if human's segment name is same as teacher's segment name??? how deal it?
	c := fucker{Name: "i am a fucker", human: human{Name: "i am human", Age: 1024, Sex: true}}
	fmt.Println("i am a fucker, i am the same Name segment in inner and human's inner, i am value =", c)
	fmt.Println("i will change the Name of c.Name...")
	c.Name = "what the fuck??? WTF"
	fmt.Println("changed of me: ", c)

	fmt.Println("i will change the Name of c.human.Name...")
	c.human.Name = "what the human fuck??? WTF"
	fmt.Println("changed of me: ", c)
}

func main() {
	assign()
	func_arg()
	anonymous_struct()
	struct_assign_and_compare()
	compose()
}
