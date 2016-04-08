package controllers

import (
	"testing"
	. "ginsite/models"
)



func TestGetTitles(t *testing.T)  {
	InitDB()
	var title []string
	DB.Table("posts").Order("pid desc").Limit(5).Pluck("title",&title)
	for i:=0;i<len(title);i++ {
			t.Logf("%v\n",title[i])
	}

}