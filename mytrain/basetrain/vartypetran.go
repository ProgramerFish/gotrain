package basetrain

import "fmt"

func t1() {
	var sum int = 17
	var count float32 = 5.0
	var mean1 float32
	var mean2 int
	//go 不支持隐式转换类型
	//mean = float32(sum)/float32(count)
	mean1 = float32(sum)/count
	mean2 = sum/int(count)
	//mean2 = float32(sum)/count
	fmt.Printf("mean 的值为: %f\n",mean1)
	fmt.Printf("mean 的值为: %d\n",mean2)
}

func t2() {
	var a int64 = 3
    var b int32
    b = int32(a)
	var aa int32 = 3
    var ab int64
    ab = int64(aa)
    fmt.Printf("b 为 :%d %d\n",b, ab)
}

func VarTypeTran() {
	t1()
	t2()

}