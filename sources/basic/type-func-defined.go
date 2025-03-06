package main

// 원형 정의
type calculator func(int, int) int

// calculator 원형 사용
func calc(f calculator, a int, b int) int {

    result := f(a, b)
    return result

}

