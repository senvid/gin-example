package controllers
 import (
 	. "ginsite/models"
 	"github.com/gin-gonic/gin"

 )


func AsideHandler(c *gin.Context)  {
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
	DB.Table("tags").Select("count(pid) as count,type as tagname").Joins("LEFT JOIN posts ON tags.tid = posts.tagid").Group("tid").Scan(&result)

	//select title from posts order by pid desc limit 5
	var title []string
	DB.Table("posts").Order("pid desc").Limit(5).Pluck("title",&title)



}
