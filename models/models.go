package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/xppcnn/gin-demo/config"
)

var db *gorm.DB

type Model struct {
	ID         string `gorm:"primary_key" json:"id" binding:"uuid"`
	CreatedOn  int    `json:"created_on"`
	ModifiedOn int    `json:"modified_on"`
	DeletedOn  int    `json:"deleted_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	Cfg := config.GetConfig()
	sec := Cfg.DataBase
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType = sec.Type
	dbName = sec.Name
	user = sec.User
	password = sec.PassWord
	host = sec.Host
	tablePrefix = sec.TablePrefix
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

}

func CloseDB() {
	defer db.Close()
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// func deleteCallback(scope *gorm.Scope){
// 	if !scope.HasError() {
// 		var extraOption string
// 		if str, ok := scope.Get(){

// 		}
// 	}
// }
