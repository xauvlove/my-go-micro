package main

import (
	"fmt"

	// 导入这个包，起空别名 _ ， 不直接使用包，但后续的程序会使用到这个包，因此需要匿名导入，使用了该包的 init() 函数
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Student struct {
	// id 会默认是主键
	Id   int
	Name string
	Age  int
}

type Teacher struct {
	gorm.Model
	Name string
	Age  int
}

func (t Teacher) String() string {
	return fmt.Sprintf("id=%v, name=%v, age=%v, deletedAt=%v\n", t.ID, t.Name, t.Age, t.DeletedAt)
}

var GlobalConn *gorm.DB

func main() {
	conn, err := gorm.Open("mysql", "root:521410@tcp(localhost:3306)/test?parseTime=true&loc=Local")
	if err != nil {
		fmt.Printf("gorm open err = %v\n", err)
	}

	GlobalConn = conn

	conn.DB().SetMaxIdleConns(50)
	conn.DB().SetMaxOpenConns(100)

	//defer conn.Close()

	// 借助 gorm 创建表，使用 AutoMigrate 函数，创建表
	//conn.SingularTable(true)
	//conn.AutoMigrate(new(Teacher))
	// conn.Error 是获取刚才的语句看看有没有错误
	//err = conn.Error
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//}
	//InsertData()
	//QueryData()
	//UpdateData()
	DeleteData()
}

func InsertData() {
	// 创建数据
	stu := new(Student)
	stu.Name = "wang wu"
	stu.Age = 19
	// 插入数据库
	GlobalConn.Create(stu)
}

func QueryData() {
	// 查询表中第一条数据 按照主键排序
	var stu = Student{}
	GlobalConn.First(&stu)
	fmt.Printf("first = %v\n", stu)

	// 查询最后一条数据，按照主键排序
	GlobalConn.Last(&stu)
	fmt.Printf("last = %v\n", stu)

	// 查询所有数据
	var stus = make([]Student, 10)
	GlobalConn.Find(&stus)
	fmt.Printf("all = %v\n", stus)

	// 只查询部分字段
	GlobalConn.Select("name").Find(&stus)
	fmt.Printf("all (only name field) = %v\n", stus)

	// 带查询条件
	GlobalConn.Select("name, age").Where("name=?", "zhang san").Find(&stus)
	fmt.Printf("all (with where) = %v\n", stus)

	// 带两个条件查询
	GlobalConn.Select("*").Where("name=?", "zhang san").Where("age=?", 22).Find(&stus)
	fmt.Printf("all (with where2) = %v\n", stus)

	// 带两个条件查询 另一种写法
	GlobalConn.Select("*").Where("name=? and age=?", "zhang san", 22).Find(&stus)
	fmt.Printf("all (with where3) = %v\n", stus)
}

func UpdateData() {
	var stu Student
	GlobalConn.First(&stu)
	fmt.Printf("基准数据 = %v\n", stu)

	// 更新数据，model(&stu) 是指定 students 表
	GlobalConn.Model(&stu).Where("id=?", stu.Id).Update("name", "zhao liu").Update("age", 100)
	// 使用 map 更新多个字段
	GlobalConn.Model(&stu).Where("id=?", stu.Id).Updates(map[string]interface{}{
		"name": "zhao liu",
		"age":  30,
	})

	// 更新后，重新拿 id 查询
	GlobalConn.Where("id=?", stu.Id).First(&stu)
	fmt.Printf("更新后数据 = %v\n", stu)
}

func DeleteData() {

	// 初始化测试数据
	{
		GlobalConn.Unscoped().Where("id>?", 0).Delete(new(Teacher))
		t := Teacher{Name: "teacher 1", Age: 33}
		t1 := Teacher{Name: "teacher 2", Age: 32}
		t2 := Teacher{Name: "teacher 3", Age: 33}
		t3 := Teacher{Name: "teacher 4", Age: 35}
		GlobalConn.Create(&t)
		GlobalConn.Create(&t1)
		GlobalConn.Create(&t2)
		GlobalConn.Create(&t3)
	}

	var teachers []Teacher
	GlobalConn.Find(&teachers)
	fmt.Printf("删除前，所有数据 = %v\n", teachers)

	GlobalConn.Where("name=?", "teacher 1").Delete(new(Teacher))
	GlobalConn.Find(&teachers)
	fmt.Printf("删除后，所有数据 = %v\n", teachers)

	GlobalConn.Unscoped().Find(&teachers)
	fmt.Printf("删除前，所有数据（包含查询已删除数据） = %v\n", teachers)

	GlobalConn.Unscoped().Where("name=?", "teacher 1").Delete(new(Teacher))
	GlobalConn.Unscoped().Find(&teachers)
	fmt.Printf("物理删除后，所有数据（包含查询已删除数据） = %v\n", teachers)
}
