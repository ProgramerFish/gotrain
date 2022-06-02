package main

import (
	"fmt"
	"github.com/wmy09527/gotrain/mytrain/basetrain"
)

/* 声明全局变量 */
var global1 int = 10

func main() {
	// this main
	/*
		    fmt.Println("This main func .")
			basetrain.Var_init_def_val()
			fmt.Println(basetrain.Numbers())
			basetrain.MyIota()
			basetrain.CalMath()
			basetrain.CalRel()
			basetrain.CalLogic()
			basetrain.CalBit()
			basetrain.LoopFor()
			basetrain.LoopRange()
			basetrain.TestFunc()
			basetrain.VarArray()
	
	basetrain.VarPoint()
	basetrain.VarPPoint()
	basetrain.VarStruct()
	
	basetrain.VarSlice()
	*/
	//basetrain.VarRange()
	//basetrain.VarMap()
	//basetrain.HashMapa()
	//basetrain.FuncRecursive()
	//basetrain.VarTypeTran()
	//basetrain.FuncIterface()
	//basetrain.InterfaceErr()
	basetrain.ThreadGoRoutine()

	//const define in func can not be used in package ,const define in head line can be used here ，e'g: const from var.go
	//private const can not use in other package
	//fmt.Println("my package const cstr:",cstr)
	//fmt.Println("my package const myconst:",basetrain.myconst)
	fmt.Println("my package const myconst:", basetrain.Myconst1)
	//c := "test"
	//fmt.Printf("format line: %s xxx %s",c,c)

	/* 声明局部变量 与全局变量同名 */
	var global1 int = 11

	fmt.Printf("结果： g = %d\n", global1)
}
