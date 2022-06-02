package basetrain

import (
	"fmt"
	"math"
)

/*
func function_name( [parameter list] ) [return_types] {
	函数体
 }
*/

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}

/*值传递  定义相互交换值的函数 */
func swapIntv(x, y int) int {
	var temp int

	temp = x /* 保存 x 的值 */
	x = y    /* 将 y 值赋给 x */
	y = temp /* 将 temp 值赋给 y*/

	return temp
}

/*引用传递  定义交换值函数*/
func swapIntp(x *int, y *int) {
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}

type cb func(int) int
type mystr string

//函数作为参数实现回调
func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}

//函数作为返回实现闭包
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

/*
 func (variable_name variable_data_type) function_name() [return_type]{
	// 函数体
 }
*/
/* 定义结构体 */
type Circle struct {
	radius float64
}

//类型的方法 ：
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

// 注意如果想要更改成功c的值，这里需要传指针
func (c *Circle) changeRadius(radius float64) {
	c.radius = radius
}

// 以下操作将不生效
//func (c Circle) changeRadius(radius float64)  {
//   c.radius = radius
//}

// 引用类型要想改变值需要传指针
func change(c *Circle, radius float64) {
	c.radius = radius
}

func TestFunc() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(a, b)

	fmt.Printf("最大值是 : %d\n", ret)

	a1, b1 := swap("Google", "Runoob")
	fmt.Println(a1, b1)

	var a2 int = 100
	var b2 int = 200

	fmt.Printf("交换前 a 的值为 : %d\n", a2)
	fmt.Printf("交换前 b 的值为 : %d\n", b2)

	/* 通过调用函数来交换值 */
	swapIntv(a2, b2)

	fmt.Printf("交换后 a 的值 : %d\n", a2)
	fmt.Printf("交换后 b 的值 : %d\n", b2)

	var a3 int = 100
	var b3 int = 200

	fmt.Printf("交换前，a 的值 : %d\n", a3)
	fmt.Printf("交换前，b 的值 : %d\n", b3)

	/* 调用 swap() 函数
	 * &a 指向 a 指针，a 变量的地址
	 * &b 指向 b 指针，b 变量的地址
	 */
	swapIntp(&a3, &b3)

	fmt.Printf("交换后，a 的值 : %d\n", a3)
	fmt.Printf("交换后，b 的值 : %d\n", b3)

	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* 使用函数 */
	fmt.Println(getSquareRoot(9))

	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，x：%d\n", x)
		return x
	})

	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	var c1 Circle
	c1.radius = 10.00
	c2 := &Circle{10.00}
	fmt.Println("圆的面积 = ", c2.getArea())
	c1.changeRadius(20.00)
	c2.changeRadius(20.00)
	fmt.Println("change c1 = 20 :", c1.radius, c2.radius)
	change(&c1, 30.00)
	change(c2, 30.00)
	fmt.Println("change c1 =30 :", c1.radius, c2.radius)

	//const define in func can not be used in package ,const define in head line can be used here ，e'g: const from var.go
	//private const can not use in other package
	//fmt.Println("my package const cstr:",cstr)
	fmt.Println("my package const myconst:", myconst)
	var mystr1 mystr = "mystr"
	fmt.Println(mystr1)

}
