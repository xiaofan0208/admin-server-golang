package models

import (
	"github.com/jinzhu/gorm"
	"strings"
)

func GetAllDatasByQuery(DB  *gorm.DB , query map[string]interface{} , filterMap map[string]interface{} , order map[string]string ,  limit int64 , offset int64) ( *gorm.DB){
	// 根据字段匹配
	for k,v := range query {
		DB = DB.Where( k + " = ?" , v )
	}

	for k,v := range filterMap {
		arr := strings.Split(  strings.TrimSpace(k) ," ")
		if len(arr) > 1 {
			key := arr[len(arr)-1]
			if key == "<>" {
				DB = DB.Where(arr[0] + " <> ?" , v )
			}else if key == "in" || key == "IN" {
				DB = DB.Where(arr[0] + " in (?)" , v )
			}else if key == "like" ||  key == "LIKE"  {
				DB = DB.Where( arr[0] + " LIKE ?" , "%" + v.(string) + "%" )
			}
		}
	}

	// 排序
	for k,v := range order {
		if v == ""{
			v = "asc"
		}
		DB = DB.Order( k + " " + v )
	}

	if limit <= 0 {
		limit = 20
	}
	DB = DB.Limit(limit).Offset(offset)

	return DB
}
