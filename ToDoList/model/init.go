package model

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Database(connstring string) {
	fmt.Println(connstring)
	db, err := gorm.Open("mysql", connstring)
	if err != nil {
		panic("Mysql数据库连接错误")
	} else {
		fmt.Println("Mysql数据库连接成功")
	}
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}

	mysqldb := db.DB()
	db.SingularTable(true)                       // 表名后不加s
	mysqldb.SetMaxIdleConns(20)                  // 设置连接池
	mysqldb.SetMaxOpenConns(100)                 //设置最大连接数
	mysqldb.SetConnMaxLifetime(time.Second * 30) //最大连接时间
	DB = db
	migration()
}
