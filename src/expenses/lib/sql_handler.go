package lib

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// SQLHandler
type SQLHandler struct {
	DB  *gorm.DB
	Err error
}

var dbConn *SQLHandler

func NewSQLHandler() *SQLHandler {
	DBMS := "mysql"
	USER := "eps_g_liu"
	PASS := "v4wedsavhoqwd"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "eps_info"

	var db *gorm.DB
	var err error
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
	}
	sqlDB := db.DB()

	sqlDB.SetConnMaxLifetime(100 * time.Second)
	sqlHandler := new(SQLHandler)
	sqlHandler.DB = db
	return sqlHandler

}

// GetDBConn ...
func GetDBConn() *SQLHandler {
	return dbConn
}

// BeginTransaction ...
func BeginTransaction() *gorm.DB {
	dbConn.DB = dbConn.DB.Begin()
	return dbConn.DB
}

// Rollback ...
func RollBack() {
	dbConn.DB.Rollback()
}

// DBOpen は DB connectionを張る。
func DBOpen() {
	dbConn = NewSQLHandler()
}

// DBClose は DB connectionを張る。
func DBClose() {
	sqlDB := dbConn.DB.DB()
	sqlDB.Close()
}
