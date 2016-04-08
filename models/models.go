package models

import (

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"


)

const (
	_DB_NAME     = "blog:blog@/ginsite?charset=utf8&parseTime=True&loc=Local"
	_DRIVER_NAME = "mysql"
)

//default table name is users
type User struct {
	Uid      uint   `gorm:"primary_key;AUTO_INCREMENT;not null"`
	Email    string `gorm:"size:30;unique_index;not null"`
	Password string `gorm:"type:varchar(50);not null"`
	Nickname string `gorm:"size:10;not null;unique"`
}
type Tag struct {
	Tid  uint   `gorm:"primary_key;AUTO_INCREMENT;not null"`
	Type string `gorm:"size:20;not null;unique"`

}
type Post struct {
	Pid       uint      `gorm:"primary_key;AUTO_INCREMENT;not null"`
	Slug      string    `gorm:"size:30;not null;unique"`
	Title     string    `gorm:"size:100;not null"`
	Content   string    `gorm:"type:text;not null"`
	Published time.Time `gorm:"not null;index"`
	Updated   time.Time `gorm:"not null;CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	Userid    uint      `gorm:"index;not null"`
	Tagid     uint      `gorm:"index;DEFAULT 0"`
}

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(_DRIVER_NAME, _DB_NAME)

	if err != nil {
		panic(err)
	}
	DB.LogMode(true)
	DB.DB().SetMaxIdleConns(5)
	DB.DB().SetMaxOpenConns(100)
	DB.DB().Ping()

	DB.AutoMigrate(&User{}, &Tag{}, &Post{})

	DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8")
	DB.Set("gorm:table_options", "ENGINE=InnoDB")

	DB.Model(&Post{}).AddForeignKey("userid", "users(uid)", "RESTRICT", "RESTRICT")
	DB.Model(&Post{}).AddForeignKey("tagid", "tags(tid)", "RESTRICT", "RESTRICT")
}












