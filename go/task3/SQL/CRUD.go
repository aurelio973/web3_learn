// 题目1：基本CRUD操作
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 创建学生模型
type Student struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"type:string"`
	Age   int    `gorm:"type:int"`
	Grade string `gorm:"type:string"`
}

func main() {
	// 初始化数据库连接
	dsn := "root:123123@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

	// 创建表
	db.AutoMigrate(&Student{})

	// 插入
	db.Create(&Student{Name: "张三", Age: 20, Grade: "三年级"})

	// 查询年龄>18的学生
	var students []Student
	db.Where("age>?", 18).Find(&students)

	// 更新年级
	db.Model(&Student{}).Where("name=?", "张三").Update("grade", "四年级")

	// 删除年龄<15的学生
	// db.Where("age<?", 15).Delete(&Student{})
}

