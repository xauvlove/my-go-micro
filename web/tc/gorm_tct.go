package main

import (
	"fmt"

	// 导入这个包，起空别名 _ ， 不直接使用包，但后续的程序会使用到这个包，因此需要匿名导入，使用了该包的 init() 函数
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "root:521410@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Printf("gorm open err = %v\n", err)
	}
	defer db.Close()

	type Student struct {
		// id 会默认是主键
		Id   int
		Name string
		Age  int
	}

	// 借助 gorm 创建表，使用 AutoMigrate 函数，创建表
	db.SingularTable(true)
	db.AutoMigrate(new(Student))
	// db.Error 是获取刚才的语句看看有没有错误
	err = db.Error
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
