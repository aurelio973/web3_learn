// 题目2：事务语句

package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 创建账户表
type Account struct {
	ID      int `gorm:"primaryKey"`
	Balance float64
}

// 创建交易记录表
type Transaction struct {
	ID            int `gorm:"primaryKey"`
	FromAccountID int
	ToAccountID   int
	Amount        float64
}

func main() {
	// 初始化数据库连接
	dsn := "root:bjw061211@tcp(127.0.0.1:3306)/bank?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

	// 创建表
	if err := db.AutoMigrate(&Account{}, &Transaction{}); err != nil {
		panic("创建表失败: " + err.Error())
	}

	initTestData(db)

	// 定义转账参数：转出账户ID、转入账户ID、转账金额
	fromID := 1
	toID := 2
	amount := 100

	// 执行转账：A向B转100元
	if err := transfer(db, fromID, toID, float64(amount)); err != nil {
		fmt.Println("转账失败:", err)
	} else {
		fmt.Println("转账成功")
	}
}

// 转账事务
func transfer(db *gorm.DB, fromID, toID int, amount float64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 查询转出账户
		var from Account
		if err := tx.First(&from, fromID).Error; err != nil {
			return fmt.Errorf("转出账户不存在: %v", err)
		}

		// 检查余额
		if from.Balance < amount {
			return fmt.Errorf("余额不足（当前: %.2f, 需转: %.2f）", from.Balance, amount)
		}

		// 扣减转出账户余额
		if err := tx.Model(&Account{}).Where("id = ?", fromID).
			Update("balance", from.Balance-amount).Error; err != nil {
			return fmt.Errorf("扣减余额失败: %v", err)
		}

		// 增加转入账户余额
		var to Account
		if err := tx.First(&to, toID).Error; err != nil {
			return fmt.Errorf("转入账户不存在: %v", err)
		}
		if err := tx.Model(&Account{}).Where("id = ?", toID).
			Update("balance", to.Balance+amount).Error; err != nil {
			return fmt.Errorf("增加余额失败: %v", err)
		}

		// 记录交易
		if err := tx.Create(&Transaction{
			FromAccountID: fromID,
			ToAccountID:   toID,
			Amount:        amount,
		}).Error; err != nil {
			return fmt.Errorf("记录交易失败: %v", err)
		}

		return nil // 成功提交事务
	})
}

// 初始化测试数据
func initTestData(db *gorm.DB) {
	// 清空测试数据
	db.Exec("DELETE FROM transactions")
	db.Exec("DELETE FROM accounts")

	// 测试
	db.Create(&Account{ID: 1, Balance: 50})
	db.Create(&Account{ID: 2, Balance: 300})
}

