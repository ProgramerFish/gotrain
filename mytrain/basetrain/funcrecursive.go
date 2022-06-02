package basetrain

import "fmt"

/*
func recursion() {
	recursion() // 函数调用自身 
 }
 
 func main() {
	recursion()
 }
 */

 func Factorial(n uint64)(result uint64) {
    if (n > 0) {
        result = n * Factorial(n-1)
        return result
    }
    return 1
}

func fibonacci(n int) int {
	if n < 2 {
	 return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
  }

func f1() {
	var i int = 15
    fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}

func f2() {
	var i int
    for i = 0; i < 10; i++ {
       fmt.Printf("%d\t", fibonacci(i))
    }
	fmt.Println("")
}

func FuncRecursive() {

	f1()
	f2()

}