package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error

	dsn := "root:bjw061211@tcp(localhost:3306)/blog_db?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接MySQL数据库
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err // 返回错误信息
	}

	// 自动迁移数据表（如果表不存在则创建）
	err = DB.AutoMigrate(
		&User{},    // 用户表
		&Post{},    // 文章表
		&Comment{}, // 评论表
	)
	return err
}

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Email    string `gorm:"unique;not null" json:"email"`
	Posts    []Post `json:"comments"`
}

type Post struct {
	gorm.Model
	Title    string    `gorm:"not null" json:"content"`
	Content  string    `gorm:"not null" json:"content"`
	UserID   uint      `json:"user_id"`
	User     User      `json:"user"`
	comments []Comment `json:"comments"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"content"`
	UserID  uint   `json:"user_id"`
	User    User   `json:"user"`
	PostID  uint   `json:"post_id"`
	Post    Post   `json:"post"`
}
