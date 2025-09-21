package main

import (
	"blog/internal/myredis"
	"blog/middlewares"
	"blog/migrate"
	"blog/validators"
	"blog/handler"
	"blog/zaplogger"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// gin程序入口文件
var (
	rdb *redis.Client // Redis 客户端实例声明（全局变量）
)

// @title Gin Web API
// @version 1.0
// @description RESTful API 文档示例
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/license/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.basic BasicAuth

func main() {
	// 初始化logger
	loggerMgr := zaplogger.InitLogger() // 初始化zap日志管理器
	zap.ReplaceGlobals(loggerMgr)       // 修复：函数名拼写错误 ReplaceGLobals -> ReplaceGlobals
	defer loggerMgr.Sync()              // 程序退出时刷新日志缓冲区
	logger := loggerMgr.Sugar()         // 创建带语法糖的日志实例
	logger.Debug("START!")              // 输出调试日志：程序启动

	// 初始化数据库
	db := migrate.InitMigrate() // 初始化数据库连接并执行自动迁移
	
	// 初始化redis 
	rdb = myredis.InitRedis() // 初始化Redis客户端并赋值给全局变量

	// 初始化gin
	r := gin.Default() // 创建 Gin 引擎实例，默认包含 Logger 和 Recovery 中间件

	// 配置Swagger.UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger.json")))

	// 配置swagger.json文件访问
	r.StaticFile("/swagger.json", "./docs/swagger.json")
	
	api := r.Group("/api")
	// 注册接口
	api.POST("/register", handler.Register(db))
	// 登录接口
	api.POST("/login", handler.Login(db, rdb))

	// 路由分版本
	v1 := r.Group("/api/v1")
	// 设置中间件 
	v1.Use(
		middlewares.LatencyLogger(),
		middlewares.CORSMiddleware(),
		middlewares.JWTAuth(rdb),
	)
	{
		// 文章Restful API接口
		v1.GET("/post", handler.QueryOnePostByTitleService(db))
		v1.GET("/post/all", handler.QueryPostListByUserId(db))
		v1.POST("/post", handler.PostCreateHandler(db))
		v1.PUT("/post", handler.UpdatePostByUserId(db))
		v1.DELETE("/post", handler.DeletePostByUserId(db))

		// 评论Restful API接口
		v1.POST("/comment", handler.CreateCommentByPostIdHandler(db))
		v1.GET("/comment", handler.QueryCommentByPostIdHandler(db))
	}

	// 注册自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("passwordReg", validators.PasswordValidator)
		if err != nil {
			return
		}
	}

	// 启动服务器放到goroutine中，避免阻塞信号监听
	go func() {
		if err := r.Run(":8080"); err != nil {
			logger.Error("服务器启动失败", zap.Error(err))
		}
	}()

	// 监听程序终止信号，进行资源释放
	listenSignal()
}

func listenSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM) // 监听中断和终止信号
	sign := <-c
	zap.S().Debug("收到退出信号", zap.String("signal", sign.String()))
	
	// 关闭Redis连接
	if rdb != nil {
		if err := rdb.Close(); err != nil {
			zap.S().Error("Redis关闭失败", zap.Error(err))
		}
	}
	
	zap.S().Debug("资源已经释放完成")
	os.Exit(0)
}

