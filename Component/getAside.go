package component

import (
	. "ginsite/models"
)

//	SELECT COUNT(pid),type FROM tags LEFT JOIN posts ON tid = tagid GROUP BY tid
//	| COUNT(pid) | type   |
//	+------------+--------+
//	|          1 | linux  |
//	|          2 | python |
//	|          5 | go     |
//	+------------+--------+

type ts []struct {
	Title string
	Slug  string
}

func GetTitle() ts {
	//select title from posts order by pid desc limit 5
	// var title []string
	// DB.Table("posts").Order("pid desc").Limit(5).Pluck("title", &title)
	var res ts
	DB.Table("posts").Order("pid desc").Limit(5).Scan(&res)

	return res
}

//
type ct []struct {
	Count   int
	Tagname string
}

func GetTag() ct {
	var res ct
	DB.Table("tags").Select("count(pid) as count,type as tagname").Joins("LEFT JOIN posts ON tags.tid = posts.tagid").Group("tid").Scan(&res)
	return res
}
