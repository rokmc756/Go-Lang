package main

func main() {

    var s []int

    if s == nil {
        println("Nil Slice")
    }

    println(len(s), cap(s)) // 모두 0
}

