package main

import "fmt"

func main(){
    var ptr *int
    var a int = 10
    ptr = &a
    fmt.Println(ptr)
    fmt.Println(*ptr)

    var ptr1 *int
    var ptr2 *float32
    var b float32 = 12.0
    ptr2 = &b

    var ret1 = ptr1==nil
    var ret2 = ptr2==nil
    fmt.Println(ret1)
    fmt.Println(ret2)
}
