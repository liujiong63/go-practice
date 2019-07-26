package main

import (
    "fmt"
    "unsafe"
)

func main(){
    const a int = 10
    const b string = "new world"
    const c bool = false

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)

    const (
        a1 = 12
	b1 = "old world"
	c1 = true
    )
    fmt.Println(a1)
    fmt.Println(b1)
    fmt.Println(c1)

    a2 := unsafe.Sizeof(a1)
    fmt.Println(a2)

    const (
        a3 = iota
	b3 = iota
	c3 = iota
    )
    fmt.Println(a3)
    fmt.Println(b3)
    fmt.Println(c3)

    const (
        a4 = iota
	b4
	c4
    )
    fmt.Println(a4, b4, c4)

    const (
        a5 = 1<<iota
	b5 = 3<<iota
	c5
	d5
    )
    fmt.Println("a5 is", a5)
    fmt.Println("b5 is", b5)
    fmt.Println("c5 is", c5)
    fmt.Println("d5 is", d5)
}
