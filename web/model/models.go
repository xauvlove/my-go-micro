package model

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID            int
	Name          string
	Password_hash string `gorm:"size:32;unique"`
	Mobile        string `gorm:"size:11;unique"`
	Real_name     string
	Id_card       string
	Avatar_url    string
	// 用户发布的房屋信息，一个用户发布多套房
	Houses []*House
	// 用户下的订单，一个用户下多个订单
	Orders []*OrderHouse
}

type House struct {
	gorm.Model
	UserId          uint
	AreaId          uint
	Title           string
	Address         string
	Room_count      int
	Acreage         int
	Price           int
	Unit            string
	Capacity        int
	Beds            string
	Deposit         int
	Min_days        int
	Max_days        int
	Order_count     int
	Index_image_url string
	// 房屋设施，家具
	Facilities []*Facility
	// 房屋图片
	Images []*HouseImage
	// 房屋订单
	Orders []*OrderHouse
}

type Area struct {
	Id   int
	Name string
	// 区域内的房屋
	House []*House
}

type Facility struct {
	Id   int
	Name string
	// 哪些房屋有这个设施
	Houses []*House
}

type HouseImage struct {
	Id  int
	Url string
	// 房屋 id
	HouseId uint
}

type OrderHouse struct {
	gorm.Model
	UserId  uint
	HouseId uint
	// 预定起始时间
	Begin_date time.Time
	// 预定结束时间
	End_date time.Time
	// 预定总天数
	Days int
	// 房屋单价
	House_price int
	// 订单金额
	Amount  int
	Status  string
	Comment string
	// 个人征信情况
	Credit bool
}

var GlobalConn *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:521410@tcp(localhost:3306)/search_house?parseTime=true&loc=Local")
	GlobalConn = db
	defer db.Close()
	if err == nil {
		db.SingularTable(true)
		db.AutoMigrate(new(User), new(House), new(Area), new(Facility), new(HouseImage), new(OrderHouse))
		return db, nil
	}
	return nil, err
}
