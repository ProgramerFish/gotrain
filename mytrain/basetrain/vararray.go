package basetrain

import "fmt"

/*

var variable_name [SIZE] variable_type

//多维数组
var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type

*/

//global define can not like below
//balance6 := [5] float32 {1000.0, 2.0, 3.4, 7.0, 50.0}
var balance6 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

func VarArray() {

	var balance1 [10]float32

	var balance2 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	balance3 := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	//var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	balance4 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	//  将索引为 1 和 3 的元素初始化
	balance5 := [5]float32{1: 2.0, 3: 7.0}

	balance5[4] = 50.0

	var salary float32 = balance1[9]

	fmt.Println(balance1, balance2, balance3, balance4, balance5, salary)

	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)

	// Step 3: 显示两行数据
	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])

	// Step 4: 访问第一个元素
	fmt.Println("第一个元素为：")
	fmt.Println(values[0][0])

	a := [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /* 第三行索引为 2 */
	}

	/*
		//range clause permits at most two iteration variables
		for i,j,v := range a {
			fmt.Println(i,j,v)
		} */

	for _, j := range a {
		for _, v := range j {
			fmt.Println(v)
		}
	}

}
