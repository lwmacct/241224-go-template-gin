package model

import (
	"gorm.io/gorm"
)

// User 用户表结构
type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"uniqueIndex;size:50" json:"username"`
	Password string `json:"password"` // 这里演示明文，生产中请存哈希
	Role     string `json:"role"`     // admin / manager / user 等
}

// AutoMigrateUser 自动建表
func AutoMigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
