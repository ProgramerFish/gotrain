package basetrain

import "fmt"

var identifier1, identifier2 int //private 
var Id1, Id2 int //public

func pri_func() int {
	return 0
}

func Var_init_def_val() {
	
    // 声明一个变量并初始化
    var a string
    fmt.Println("string: ",a)

    // 没有初始化就为零值
    var b int
    fmt.Println("int: ",b)

    // bool 零值为 false
    var c bool
    fmt.Println("bool: ",c)

    //def value nil
	var d *int
	fmt.Println("*int: ",d)

	//auto type 
	var e = true
	fmt.Println("auto type: ",e)

	// new var
	//var intVal int  ,error syntax
	intVal := 1 
	fmt.Println("intVal: ",intVal)

	//多声明变量
	vn1, vn2, vn3 := a, b, c
	fmt.Println("muti var: ",vn1, vn2, vn3)

	//指针类型
	vp := &intVal
	fmt.Println("step 1 intVal:",intVal)
	*vp += 1
	fmt.Println("step 2 vp, intVal :",*vp, intVal)

	//const val 
	const cstr = "abc"
	fmt.Println("cstr: ",cstr)

	// enum 
	const (
		Unknown = 0
		Female = 1
		Male = 2
	)


}

func MyIota() {
	fmt.Println("iota --- ")
	//iota 

	/* 
	const (
    a = iota
    b = iota
    c = iota
    ) 

	const (
		a = iota
		b
		c
	) */

	const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		ii          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,ii)

	const (
		i=1<<iota
		j=3<<iota
		k
		l
	)
	fmt.Println(i,j,k,l)

	fmt.Println(" --- iota ")
}

func Numbers() (int,int,string) {
	a , b , c := 1 , 2 , "str"
	return a,b,c
  }
