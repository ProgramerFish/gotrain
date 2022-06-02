package basetrain

import (
	"errors"
	"fmt"
)

/*

//golang error interface

type error interface {
    Error() string
}

*/


// 定义一个 DivideError 结构
type DivideError struct {
    dividee int
    divider int
}

// 实现 `error` 接口
func (de *DivideError) Error() string {
    strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
    return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
    if varDivider == 0 {
            dData := DivideError{
                    dividee: varDividee,
                    divider: varDivider,
            }
            errorMsg = dData.Error()
            return
    } else {
            return varDividee / varDivider, ""
    }

}






func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // 实现
	fmt.Println("nothing to do .")
	return 0,nil
}



func if1() {
	_, err:= Sqrt(-1)

	if err != nil {
	   fmt.Println(err)
	}
}

func if2() {
	// 正常情况
    if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
}
// 当除数为零的时候会返回错误信息
if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ",errorMsg)
}

}

func if3() {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
    a := 101
    b := fmt.Sprintf(strFormat, a)
	fmt.Println(b)

}

func if4() (a int, s string) {
	//s = "nothing to do 1"
	return  //默认返回 a,s
}

func if5() (int, string) {
	//返回参数未命名需要return 具体值
	return 0,"nothing to do 2"
}

func InterfaceErr() {
	if1()
	//if2()
	//if3()
	fmt.Println(if4())
	fmt.Println(if5())

}