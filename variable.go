package main

import "fmt"

func main() {
    var a string = "i am a string"
    var b, c int = 10, 11
    fmt.Println(a)
    fmt.Println(b, c)

    var a1 = "another string"
    var b1 int
    var c1 bool
    fmt.Println(a1, b1, c1)

    a2 := "string again"
    fmt.Println(a2)

    var a3, b3 = 12, "b3 is a string"
    var (
	    a4 int
	    b4 string
    )
    fmt.Println(a3, b3)
    fmt.Println(a4, b4)

    _, a5 := 13, "omg"
    // fmt.Println(_)
    fmt.Println(a5)

    _, a6, b6 := numbers()
    fmt.Println(a6, b6)
}

func numbers()(int, int, string){
    a, b, c := 1, 2, "haha"
    return a, b, c
}
