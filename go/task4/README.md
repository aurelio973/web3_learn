# go-blog

#### 安装依赖

```
1.Web 框架与接口文档
github.com/gin-gonic/gin
搭建 HTTP 服务、定义路由、处理请求响应、支持中间件

github.com/swaggo/gin-swagger + github.com/swaggo/files
自动生成 Swagger 可视化接口文档，支持在线调试

2.数据存储
gorm.io/gorm
ORM 框架，定义数据模型、执行数据库 CRUD 操作、支持关联查询

gorm.io/driver/mysql
GORM 的 MySQL 驱动，实现与 MySQL 数据库的连接

github.com/go-redis/redis/v8
Redis 客户端，缓存用户 Token、实现 Token 黑名单

3.用户认证与配置
github.com/golang-jwt/jwt/v5
生成和解析 JWT Token，实现用户登录认证

github.com/joho/godotenv
读取 .env 配置文件，管理 JWT 密钥、数据库密码等敏感信息

4.参数校验与日志

github.com/go-playground/validator/v10
校验请求参数合法性（必填项、格式校验等）

go.uber.org/zap
输出调试、错误、业务日志

Go 标准库（fmt/net/http/strings/context/log）
字符串处理、HTTP 状态码定义、上下文管理、基础日志记录

```