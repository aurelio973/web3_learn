package mysqldb

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库
func InitDB() (*gorm.DB, error) {
	// 修复1：明确指定.env文件的绝对路径（避免路径问题）
	projectRoot, err := os.Getwd() // 获取当前工作目录（项目根目录）
	if err != nil {
		return nil, fmt.Errorf("获取项目路径失败: %w", err)
	}
	envPath := filepath.Join(projectRoot, ".env") // 拼接.env的绝对路径
	
	// 加载.env文件，并处理错误
	if err := godotenv.Load(envPath); err != nil {
		return nil, fmt.Errorf("加载.env文件失败（路径: %s）: %w", envPath, err)
	}

	var (
		dbHost = os.Getenv("DB_HOST") 
		dbPort = os.Getenv("DB_PORT") 
		dbUser = os.Getenv("DB_USER") 
		dbPass = os.Getenv("DB_PASS")
		dbName = os.Getenv("DB_NAME") 
	)

	// 打印读取到的配置（调试用，确认是否正确读取）
	fmt.Printf("读取到的数据库配置：\n")
	fmt.Printf("DB_HOST: %q, DB_PORT: %q, DB_USER: %q, DB_PASS: %q, DB_NAME: %q\n",
		dbHost, dbPort, dbUser, dbPass, dbName)

	// 检查必要配置是否为空
	if dbHost == "" || dbPort == "" || dbUser == "" || dbName == "" {
		return nil, fmt.Errorf("数据库配置不完整，请检查.env文件")
	}

	// 拼接数据库连接串（DNS）
	dnsUrl := fmt.Sprintf("%s:%s", dbHost, dbPort) // 主机:端口
	DNS := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dnsUrl, dbName)
	fmt.Println("数据库连接串（DNS）:", DNS)

	// 连接数据库
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       DNS,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("数据库连接失败（DNS: %s）: %w", DNS, err)
	}

	return db, nil
}

