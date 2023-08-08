package models

import (
	"database/sql"
	"fmt"
	"go-gin/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func Setup() {

	// https: //gorm.io/docs/connecting_to_the_database.html
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		setting.DataBaseSetting.User,
		setting.DataBaseSetting.Password,
		setting.DataBaseSetting.Host,
		setting.DataBaseSetting.Port,
		setting.DataBaseSetting.DBname)
	sqlDB, sqlOpenErr := sql.Open("mysql", dsn)
	if sqlOpenErr != nil {
		log.Fatalf("models.Setup.sqlOpenerr err: %v", sqlOpenErr)
	}
	gormDB, gormOpenErr := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})
	if gormOpenErr != nil {
		log.Fatalf("models.Setup.gormOpenErr err: %v", gormOpenErr)
		sqlDB.Close() //unnecessary
	}
	// todo 添加callback
	gormDB.Callback().Create().Replace("gorm:update_time_stamp", func(db *gorm.DB) {

	})
	gormDB.Callback().Update().Replace("gorm:update_time_stamp", func(db *gorm.DB) {

	})
	gormDB.Callback().Delete().Replace("gorm:delete", func(db *gorm.DB) {

	})
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

}
