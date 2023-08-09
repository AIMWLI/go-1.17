package models

type Auth struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	if username != "" && password != "" {
		db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	}
	if auth.Id > 0 {
		return true
	}
	return false
}
