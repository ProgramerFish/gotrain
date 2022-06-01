package main

import (
    "fmt"
	"github.com/wmy09527/gotrain/mytrain/basetrain"
)

func main() {
    // this main 
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
	//c := "test"
	//fmt.Printf("format line: %s xxx %s",c,c)
}
