package basetrain

import "fmt"

/*

type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}

variable_name := structure_variable_type {value1, value2...valuen}
或
variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}

*/

type Books struct {
	title string
	author string
	subject string
	book_id int
 }

func VarStruct() {
	    // 创建一个新的结构体
		//err example:
		//fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程"})
		fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

		// 也可以使用 key => value 格式
		fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})
	
		// 忽略的字段为 0 或 空
	   fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
	   book := &Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407}
	   (*book).book_id = 123456
	   book.book_id = 123456
	   fmt.Println(book)
 
}