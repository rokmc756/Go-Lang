package main

import "fmt"

func main() {

    s := []int{0, 1, 2, 3, 4, 5}

    s = s[2:5]      // 2, 3, 4

    s = s[1:]       // 3, 4

    fmt.Println(s)  // 3, 4 출력

}

