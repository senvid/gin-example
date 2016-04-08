package models

import (
	// "fmt"

	//"github.com/jinzhu/gorm"
	"testing"
)

//must be enabled
func TestInitDB(t *testing.T) {
	defer func() {
		InitDB()
		if err := recover(); err != nil {
			t.Error("init err ")
		}
	}()
}

//
//func TestInsert(t *testing.T) {
//	u := User{
//		Email:    "test@sina.com",
//		Password: "test",
//		Nickname: "test",
//	}
//
//	//db, _ := gorm.Open("mysql", "blog:blog@/ginsite?charset=utf8&parseTime=True&loc=Local")
//	db := DB
//
//	first := db.NewRecord(u)
//	t.Logf("%v", first)
//	db.Create(&u)
//	second := db.NewRecord(u)
//	if second == false {
//		t.Log("insert sucess..")
//	}
//	db.Delete(&u)
//}

type Et struct {
	ID   uint
	UtID int    `gorm:"index"`
	Name string `gorm:"unique"`
}
type Ut struct {
	ID   uint
	Name string
	Ets  []Et
}
//
//func TestSelect(t *testing.T) {
//
//	// var db *gorm.DB
//	// var err error
//	db := DB
//	// db.DropTable(&Ut{}, &Et{})
//	db.AutoMigrate(&Ut{}, &Et{})
//
//	// u := Ut{
//	// 	Name: "s1",
//	// 	Ets: []Et{
//	// 		Et{Name: "ss1@t.com"},
//	// 	},
//	// }
//	// u.Ets = append(u.Ets, Et{Name: "ss2@t.com"})
//	// db.Create(&u)
//	var tus []Et
//	rows, e := db.Select("*").Find(&tus).Rows()
//	if e != nil {
//		t.Error("fail")
//	}
//
//	for rows.Next() {
//
//		rows.Scan(&tus)
//
//		t.Logf("%v", tus)
//	}
//	// db.Find(&Ut{}).Scan(&tus)
//	var ct uint
//	db.Model(&Et{}).Count(&ct)
//	// db.Model(&Ut{}).Order("id desc").Find(&Ut{}).Scan(&tt)
//	t.Logf("%v\n---一共%v条", tus, ct)
//	// t.Logf("*****%v*****---%v********", tt, rows)
//	// db.Delete(&u)
//}

//func TestGetTagid(t *testing.T) {
//	db := DB
//	db.LogMode(true)
//
//	var ets Et
//	//db.Where("name = ?","sun").Find(&ets)
//
//	//db.Where("name = ? ","sun").Model(&ets).Scan(&ets)
//	row := db.Where("name = ? ", "test").Model(&ets).Row()
//	e := row.Scan(&ets.ID, &ets.UtID, &ets.Name)
//	if e!=nil {
//		t.Fail()
//	}else {
//		t.Logf("id: %v   name: %v   utid: %v", ets.ID, ets.Name, ets.UtID)
//	}
//}

func TestGetCountAndTagtype(t *testing.T) {

	//	SELECT COUNT(pid),type FROM tags LEFT JOIN posts ON tid = tagid GROUP BY tid
	//	| COUNT(pid) | type   |
	//	+------------+--------+
	//	|          1 | linux  |
	//	|          2 | python |
	//	|          5 | go     |
	//	+------------+--------+
	var result []struct{
		Count int
		Tagname string
	}
	DB.LogMode(true)
	//rows,err :=DB.Table("tags").Select("count(posts.pid) as count,tags.type as tagtype").Joins("LEFT JOIN posts ON tags.tid = posts.tagid").Group("tags.tid").Rows()
	DB.Table("tags").Select("count(pid) as count,type as tagname").Joins("LEFT JOIN posts ON tags.tid = posts.tagid").Group("tid").Scan(&result)

	//if err!=nil {
	//	t.Log("no found..")
	//}
	//for rows.Next(){
	//	rows.Scan(result)
	//}
	t.Logf("%v  %v",result,len(result))

}
