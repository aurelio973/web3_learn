package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 用户模型
type User struct {
	gorm.Model
	Username string `gorm:"not null;unique"`
	Email    string `gorm:"not null;unique"`
	Passward string `gorm:"not null"`
	Posts    []Post `gorm:"foreignKey:UserID"`
}

// 文章模型
type Post struct {
	gorm.Model
	Title    string    `gorm:"not null"`
	Content  string    `gorm:"not null`
	UserID   uint      `gorm:"not null"`
	User     User      `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:PostID"`
}

// 评论模型
type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	PostID  uint   `gorm:"not null"`
	Post    Post   `gorm:"foreignKey:PostID"`
	UserID  uint   `gorm:"not null"`
	User    User   `gorm:"foreignKey:UserID"`
}

func main() {
	// 连接数据库
	dsn := "root:bjw061211@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}
	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		panic("创建表失败: " + err.Error())
	}

	println("数据表创建成功！")

}

