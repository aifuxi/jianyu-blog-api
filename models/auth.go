package models

import "fmt"

type Auth struct {
	ID       int    `gorm:"primaryKey;" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth

	fmt.Printf("数据库引用对象db: %v\n", db)

	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)

	return auth.ID > 0
}
