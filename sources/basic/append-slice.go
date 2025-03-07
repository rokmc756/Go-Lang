package main

import "fmt"

func main() {

    s := []int{0, 1}

    // 하나 확장
    s = append(s, 2)       // 0, 1, 2

    // 복수 요소들 확장
    s = append(s, 3, 4, 5) // 0,1,2,3,4,5

    fmt.Println(s)

}

