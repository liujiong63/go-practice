package main

import "fmt"

func max(a1, a2 int) int {
    var result int
    if (a1 > a2){
        result = a1
    } else {
        result = a2
    }
    return result
}

func main(){
    var ret = max(10, 11)
    fmt.Println(ret)
}
