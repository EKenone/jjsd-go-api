package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jjsd-go-api/api/mini/conf"
	"log"
	"sync"
	"time"
)

var (
	link     *gorm.DB
	linkOnce sync.Once
)

func newDbLink() *gorm.DB {
	db, err := gorm.Open(mysql.Open(getDsn()), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		log.Println(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(3)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour * 2)

	return db
}

func DbLink() *gorm.DB {
	linkOnce.Do(func() {
		link = newDbLink()
	})
	return link
}

func getDsn() (dsn string) {
	dsn = conf.Conf.Mysql.User + ":" + conf.Conf.Mysql.Pass + "@tcp(" + conf.Conf.Mysql.Host + ":" + conf.Conf.Mysql.Port + ")/jjsd?charset=utf8&loc=Asia%2FShanghai&parseTime=true"
	return
}
