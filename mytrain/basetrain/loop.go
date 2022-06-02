package basetrain

import "fmt"

func LoopFor() {
	sum := 0
	//for init; condition; post { }
	//for condition { }
	//for { }
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	sum = 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	sum = 1
	// 这样写也可以，更像 While 语句形式
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)
	//unlimit loop
	/*
			sum = 0
		    for {
		        sum++ // 无限循环下去
		    }
		    fmt.Println(sum) // 无法输出
	*/

}

func LoopRange() {
	/*
		for key, value := range oldMap {
			newMap[key] = value
		}

		for key := range oldMap

		for _, value := range oldMap
	*/

	strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	// 读取 key 和 value
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	// 读取 key
	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}

	// 读取 value
	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}

}
