package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func main() {
	// 连接数据库
	dsn := "root:123123@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True"

	// 连接数据库
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("数据库连接失败: %v\n", err)
		return
	}
	defer db.Close()

	// 查询价格大于50元的书籍
	minPrice := 50.0
	expensiveBooks, err := getBooksPriceGreaterThan(db, minPrice)
	if err != nil {
		fmt.Printf("查询失败: %v\n", err)
		return
	}

	// 打印查询结果
	fmt.Printf("价格大于%.2f元的书籍有：\n", minPrice)
	for _, book := range expensiveBooks {
		fmt.Printf("ID: %d, 书名: %s, 作者: %s, 价格: %.2f元\n",
			book.ID, book.Title, book.Author, book.Price)
	}
}

// 参数使用float64类型，与结构体Price字段类型一致，确保类型安全
func getBooksPriceGreaterThan(db *sqlx.DB, minPrice float64) ([]Book, error) {
	var books []Book
	// 使用参数化查询，避免SQL注入，同时保证类型匹配
	query := "SELECT id, title, author, price FROM books WHERE price > ?"
	err := db.Select(&books, query, minPrice)
	if err != nil {
		return nil, fmt.Errorf("查询执行失败: %w", err)
	}
	return books, nil
}

func Test(db *sqlx.DB) {
	var count int
	db.Get(&count, "SELECT COUNT(*) FROM books")
	if count == 0 {
		testBooks := []Book{
			{Title: "Go编程实战", Author: "张三", Price: 69.9},
			{Title: "Python入门", Author: "李四", Price: 45.5},
			{Title: "数据库原理", Author: "王五", Price: 79.0},
			{Title: "算法导论", Author: "赵六", Price: 128.0},
			{Title: "网络编程", Author: "孙七", Price: 55.0},
		}
		// 批量插入
		for _, b := range testBooks {
			db.NamedExec("INSERT INTO books (title, author, price) VALUES (:title, :author, :price)", b)
		}
	}
}

