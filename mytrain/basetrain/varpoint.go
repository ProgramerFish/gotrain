package basetrain

import "fmt"

const MAX int = 3

func VarPoint() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)
	fmt.Printf("ip 变量储存的指针地址: %x\n", &ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	a1 := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int
	var ptr1 *[]int
	ptr1 = &a1

	for i = 0; i < MAX; i++ {
		ptr[i] = &a1[i] /* 整数地址赋值给指针数组 */
	}
	fmt.Println(*ptr1)
	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d %d\n", i, *ptr[i], (*ptr1)[i])
	}

	for _, v := range *ptr1 {
		fmt.Println(v)
	}

}

func VarPPoint() {

	var a int
	var ptr *int
	var pptr **int
	//var pptr1 **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
	sw1, sw2 := 1, 2
	Swappoint1(&sw1, &sw2)
	fmt.Println(sw1, sw2)

}

func Swappoint1(x *int, y *int) {
	*x, *y = *y, *x

	var a int = 100
	var b int = 200
	fmt.Printf("交换前 a b 的值 : %d %d\n", a, b)
	a, b = b, a
	fmt.Printf("交换后 a b 的值 : %d %d\n", a, b)
}
