package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 题目1：使用SQL扩展库进行查询
// 创建员工表
type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

func main() {
	// 连接数据库
	dsn := "root:123123@tcp(127.0.0.1:3306)/company?charset=utf8mb4&parseTime=True"

	// 连接数据库
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("数据库连接失败: %v\n", err)
		return
	}
	defer db.Close()

	techEmployees, err := getTechDepartmentEmployees(db)
	if err != nil {
		fmt.Printf("查询技术部员工失败: %v\n", err)
	} else {
		fmt.Println("技术部员工列表:")
		for _, emp := range techEmployees {
			fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 工资: %d\n",
				emp.ID, emp.Name, emp.Department, emp.Salary)
		}
	}

	// 2. 查询工资最高的员工
	topSalaryEmp, err := getHighestSalaryEmployee(db)
	if err != nil {
		fmt.Printf("查询最高工资员工失败: %v\n", err)
	} else {
		fmt.Printf("\n工资最高的员工: ID: %d, 姓名: %s, 部门: %s, 工资: %d\n",
			topSalaryEmp.ID, topSalaryEmp.Name, topSalaryEmp.Department, topSalaryEmp.Salary)
	}
}

// 查询技术部所有员工
func getTechDepartmentEmployees(db *sqlx.DB) ([]Employee, error) {
	var employees []Employee
	err := db.Select(&employees, "SELECT * FROM employees WHERE department=?", "技术部")
	if err != nil {
		return nil, err
	}
	return employees, nil
}

// 查询工资最高的员工
func getHighestSalaryEmployee(db *sqlx.DB) (Employee, error) {
	var emp Employee
	err := db.Get(&emp, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		return Employee{}, err
	}
	return emp, nil
}
