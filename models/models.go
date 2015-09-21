package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

const (
	_DB_NAME     = "blog:blog@/blog?charset=utf8&parseTime=True&loc=Local"
	_DRIVER_NAME = "mysql"
)

type Users struct {
	Uid      uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	Email    string `sql:"size:50;not null;unique"`
	Password string `sql:"size:50;not null"`
	Nickname string `sql:"size:30;not null;unique"`
}

type Posts struct {
	Id        uint64    `sql:"primary key:AUTO_INCREMENT"`
	Slug      string    `sql:"size:50;not null;unique"`
	Title     string    `sql:"size:100;not null"`
	Content   string    `sql:"size:5000;not null"`
	Published time.Time `sql:"not null;index"`
	Updated   time.Time `sql:"NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
type Tags struct {
	Id       uint8  `sql:"primary key:AUTO_INCREMENT"`
	Parentid uint8  `sql:"default null"`
	Type     string `sql:"size:20;not null"`
}

var DB gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(_DRIVER_NAME, _DB_NAME)

	if err != nil {
		panic(err)
	}
	DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8")
	DB.Set("gorm:table_options", "ENGINE=InnoDB")
	DB.AutoMigrate(&Users{}, &Posts{}, &Tags{})
	DB.LogMode(true)
}

// type Result struct {
// 	Uid   uint
// 	Email string
// }

// var Res Result

// func GetPosts() Result {

// 	// var db *gorm.DB
// 	// var err error
// 	db, _ := gorm.Open("mysql", "blog:blog@/blog?charset=utf8&parseTime=True&loc=Local")
// 	// db.DB()
// 	// db.DB().Ping()
// 	// db.DB().SetMaxIdleConns(10)
// 	// db.DB().SetMaxOpenConns(100)
// 	// db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Users{})
// 	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Users{})
// 	db.AutoMigrate(&Users{})
// 	u := Users{Email: "test@test.com", Password: "test", Nickname: "test"}
// 	db.NewRecord(u)
// 	db.Create(&u)
// 	db.Raw("SELECT uid,email FROM users where uid=1").Scan(&Res)

// 	return Res
// }
