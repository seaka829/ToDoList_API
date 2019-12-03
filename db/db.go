package db

import (
	"../entity"
	"github.com/jinzhu/gorm"
)

const (
	dbms     = "mysql"
	user     = "root"
	pass     = "root"
	protocol = "tcp(127.0.0.1:3306)"
	dbnm     = "todo"
	connect  = user + ":" + pass + "@" + protocol + "/" + dbnm
)

var (
	db  *gorm.DB
	err error
)

// Init ... DB接続
func Init() {
	db, err = gorm.Open(dbms, connect)
	if err != nil {
		panic(err.Error())
	}
	autoMigrate()
}

// GetDB ... DBオブジェクトの取得
func GetDB() *gorm.DB {
	return db
}

// Close ... DB切断
func Close() {
	err = db.Close()
	if err != nil {
		panic(err)
	}
}

// autoMigrate ... オートマイグレーション
func autoMigrate() {
	db.AutoMigrate(&entity.Task{})
}
