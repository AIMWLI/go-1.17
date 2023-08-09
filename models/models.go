package models

import (
	"database/sql"
	"fmt"
	"go-gin/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

type Model struct {
	ID         int `json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

var db *gorm.DB

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
	var gormOpenErr error
	//db, gormOpenErr = gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})
	db, gormOpenErr = gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{
		// https://gorm.io/zh_CN/docs/gorm_config.html
		Logger: logger.Default.LogMode(logger.Info),

		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "t_",                              // table name prefix, table for `User` would be `t_users`
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			//NoLowerCase:   true, // skip the snake_casing of names
			//NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
	})
	if gormOpenErr != nil {
		log.Fatalf("models.Setup.gormOpenErr err: %v", gormOpenErr)
		sqlDB.Close() //unnecessary
	}
	// todo 添加callback
	/*
		db.Callback().Create().Replace("gorm:update_time_stamp", func(db *gorm.DB) {

		})
		db.Callback().Update().Replace("gorm:update_time_stamp", func(db *gorm.DB) {

		})
		db.Callback().Delete().Replace("gorm:delete", func(db *gorm.DB) {

		})
	*/
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

}
