package models

import (
	"github.com/go-playground/validator/v10"
	"log"
)

type User struct {
	Model
	Id int64 `gorm:"index" json:"id"`
	// https://matthung0807.blogspot.com/2021/12/go-playground-validator-exmple.html
	Name       string `json:"name" validate:"required"`
	CreatedBy  string `json:"created_by" validate:"omitempty"`
	ModifiedBy string `json:"modified_by" json:"modified_by,omitempty"`
	CreatedOn  string `json:"created_on"`
	ModifiedOn string `json:"modified_on"`
	DeletedOn  string `json:"deleted_on"`
}

func SelectPage(pageNum int, pageSize int, maps interface{}) (users []User) {
	db.Offset(pageNum).Limit(pageSize).Find(&users)
	return
}

func Save(user *User) bool {
	err := validator.New().Struct(user)
	if err != nil {
		log.Printf("[WARN] Save.validate err: %v", err)
		return false
	}
	//db.Create(u)
	db.Save(user)
	//db.Model(&user).Update("CreatedOn", time.Now())
	return true
}
