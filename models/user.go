package models

import (
	"tll/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"` //后面考虑更改，名字可能重复，不能unique
	Password string
	Role     int `gorm:"check:role IN (0, 1, 2)"` // 添加一个检查约束，0为管理员，1为教师，2为学生
}

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}
func FindUser(db *gorm.DB, user *User) error {
	return db.First(user, user.ID).Error
}

func FindUserIDByUsername(db *gorm.DB, username string) (uint, error) {
	var user struct {
		ID uint
	}
	err := db.Model(&User{}).
		Select("id").
		Where("username = ?", username).
		First(&user).Error

	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func GiveTeacherRole(db *gorm.DB, username string) error {
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return err 
	}
	return db.Model(&user).Update("role", 1).Error
}


func CreateAdminIfNotExists(db *gorm.DB) error {
	var count int64
	db.Model(&User{}).Count(&count)
	if count == 0 {
		admin := User{
			Username: "admin",
			Password: utils.HashPasswd("tll123"), // hash 密码
			Role:     0,                          // 管理员
		}
		return db.Create(&admin).Error
	}
	return nil
}

func InitUser(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
