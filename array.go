package main

import "fmt"

func main(){
    var a1 [5] int
    var a2 [10] float32
    var a3 = [5]int{1,2,3,4,5}
    var a4 = []int{6,7,8,9,10}
    fmt.Println(a1)
    fmt.Println(a2)
    fmt.Println(a3)
    fmt.Println(a4)

    var b1 = a1[1]
    b2 := a2[2]
    b3 := a3[3]
    var b4 = a4[4]
    fmt.Println(b1)
    fmt.Println(b2)
    fmt.Println(b3)
    fmt.Println(b4)

    for i:=0; i<len(a1); i++ {
        a1[i] = i+1
    }
    fmt.Println(a1)
}
