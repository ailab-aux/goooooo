// string type.

package main
import "fmt"

func test() {
    //1.string default value is ""
    //2.you can get one char by s[i]
    //3.you can't get one char's pointer by &s[i]
    //4.you can't modify the byte array.
    //5.the tail of string is not NULL.

    //runtime.h
    // struct String
    // {
    //     byte* str;
    //     intgo len;
    //};

    s := "abc"
    println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)

    //`` for no transfer string, is raw string.
    ss := `a
    b\r\nx00
    c`

    println(ss)

    // concat two string use +,  + must be in the end of the first line.
    str := "hello, " +
    "world"
    // the following will be compile error
    //str2 := "hwllo, "
    //+ "world"
    println(str)

    // like python, you can use range[begin:end]
    s1 := str[:5]
    s2 := str[7:]
    s3 := str[1:5]
    println(&str)
    println(&s1, s1)
    println(&s2, s2)
    println(&s3, s3)


    // '' char const variable is Unicode Code Point, type is rune
    // rune is int32's rename.it support \uffff, \U7fffffff, \xff format.
    // rune type, UCS-4
    fmt.Printf("%T\n", 'a')

    var c1, c2 rune = '\u6211', '们'
    println(c1 == '我', string(c2) == "\xe4\xbb\xac")

    // if you want to modify string, you should cast to []rune or []bute first, modify over, cast to string.
    // cast will re-alloc memory, and copy byte array.

    a := "abcd"
    bs := []byte(s)

    bs[1] = 'B'
    a1 := string(bs)
    println(&a, &bs, &a1)
    println(a, bs, a1)

    u := "电脑"
    us := []rune(u)
    us[1] = '话'
    u1 := string(us)
    println(&u, &us, &u1)
    println(u, us, u1)

    // for loop string.
    astr := "abc汉子"
    for i := 0; i < len(astr); i++ {
        fmt.Printf("%c,", astr[i])
    }
    fmt.Println()

    for _, r := range astr {
        fmt.Printf("%c,", r)
    }
    fmt.Println()
}

func main() {
    test()
}
