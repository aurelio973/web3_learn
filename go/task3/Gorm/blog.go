package main

import (
	"fmt"
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
	PostCount   int `gorm:"default:0"`
}

// 文章模型
type Post struct {
	gorm.Model
	Title        string    `gorm:"not null"`
	Content      string    `gorm:"not null`
	UserID         uint    `gorm:"not null"`
	User           User    `gorm:"foreignKey:UserID"`
	Comments  []Comment    `gorm:"foreignKey:PostID"`
	CommentState string    `gorm:"default:'无评论'"`
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

	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	// 测试钩子函数
	testHooks(db)

}
// 创建测试数据（仅首次运行时需要）
	initTestData(db)

	// 1. 查询ID为1的用户发布的所有文章及对应的评论
	userPosts, err := getUserPostsWithComments(db, 1)
	if err != nil {
		fmt.Printf("查询用户文章失败: %v\n", err)
	} else {
		fmt.Println("===== 用户发布的文章及评论 =====")
		for _, post := range userPosts {
			fmt.Printf("\n文章ID: %d, 标题: %s\n", post.ID, post.Title)
			fmt.Println("评论列表:")
			for _, comment := range post.Comments {
				fmt.Printf("- 评论ID %d: %s (评论者ID: %d)\n", comment.ID, comment.Content, comment.UserID)
			}
		}
	}

	// 2. 查询评论数量最多的文章
	topPost, err := getPostWithMostComments(db)
	if err != nil {
		fmt.Printf("查询评论最多的文章失败: %v\n", err)
	} else {
		fmt.Printf("\n===== 评论最多的文章 =====")
		fmt.Printf("\n文章ID: %d, 标题: %s, 评论数量: %d\n",
			topPost.ID, topPost.Title, len(topPost.Comments))
	}
}

func getUserPostsWithComments(db *gorm.DB, userID uint) ([]Post, error) {
	var posts []Post
	err := db.Where("user_id = ?", userID).Preload("Comments").Find(&posts).Error
	return posts, err
}

func getPostWithMostComments(db *gorm.DB) (Post, error) {
	var post Post
	subQuery := db.Model(&Comment{}).Select("post_id, COUNT(*) as comment_count").Group("post_id").Order("comment_count DESC").Limit(1)
	err := db.Model(&Post{}).Joins("JOIN (?) as sub ON posts.id = sub.post_id", subQuery).Preload("Comments").First(&post).Error

	return post, err
}

func initTestData(db *gorm.DB) {
	// 清空现有数据
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Comment{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Post{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})

	// 创建测试用户
	// 创建测试用户（确保ID被正确赋值）
	var user User
	if err := db.Create(&User{
		Username: "test_user",
		Email:    "test@example.com",
		Password: "123456",
	}).Scan(&user).Error; err != nil {
		panic("创建用户失败: " + err.Error())
	}

	// 创建测试文章
	post1 := Post{Title: "第一篇文章", Content: "这是第一篇文章的内容", UserID: user.ID}
	post2 := Post{Title: "第二篇文章", Content: "这是第二篇文章的内容", UserID: user.ID}
	db.Create(&post1)
	db.Create(&post2)

	// 创建测试评论（给第一篇文章多加点评论，使其成为评论最多的文章）
	db.Create(&Comment{Content: "好文章！", PostID: post1.ID, UserID: user.ID})
	db.Create(&Comment{Content: "受益匪浅", PostID: post1.ID, UserID: user.ID})
	db.Create(&Comment{Content: "期待更新", PostID: post2.ID, UserID: user.ID})
}


// Post模型的创建前钩子：自动更新用户文章数量
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	// 1. 先查询用户当前的文章数量
	var user User
	if err := tx.First(&user, p.UserID).Error; err != nil {
		return fmt.Errorf("查询用户失败: %v", err)
	}

	// 2. 更新用户的文章数量
	if err := tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", user.PostCount+1).Error; err != nil {
		return fmt.Errorf("更新文章数量失败: %v", err)
	}
	return nil
}

// Comment模型的删除后钩子：检查文章评论数量并更新状态
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	// 1. 查询当前文章的剩余评论数量
	var commentCount int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount).Error; err != nil {
		return fmt.Errorf("查询评论数量失败: %v", err)
	}
	// 2. 根据评论数量更新文章状态
	state := "有评论"
	if commentCount == 0 {
		state = "无评论"
	}

	if err := tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_state", state).Error; err != nil {
		return fmt.Errorf("更新评论状态失败: %v", err)
	}
	return nil
}

// 测试钩子函数功能
func testHooks(db *gorm.DB) {
	// 清空数据
	db.Unscoped().Delete(&Comment{})
	db.Unscoped().Delete(&Post{})
	db.Unscoped().Delete(&User{})

	// 1. 创建用户
	var user User
	db.Create(&User{Username: "test", Email: "test@example.com", Password: "123"})
	db.First(&user, "username = ?", "test")
	fmt.Printf("初始用户文章数: %d\n", user.PostCount) // 0

	// 2. 创建文章（触发Post的BeforeCreate钩子）
	post := Post{Title: "测试文章", Content: "测试内容", UserID: user.ID}
	db.Create(&post)
	
	// 验证用户文章数是否增加
	db.First(&user, user.ID)
	fmt.Printf("创建文章后用户文章数: %d\n", user.PostCount) // 1

	// 3. 创建评论
	comment := Comment{Content: "测试评论", PostID: post.ID, UserID: user.ID}
	db.Create(&comment)
	
	// 验证文章初始评论状态
	var postAfterComment Post
	db.First(&postAfterComment, post.ID)
	fmt.Printf("创建评论后文章状态: %s\n", postAfterComment.CommentState) // 有评论

	// 4. 删除评论（触发Comment的AfterDelete钩子）
	db.Delete(&comment)
	
	// 验证文章评论状态是否更新
	db.First(&postAfterComment, post.ID)
	fmt.Printf("删除评论后文章状态: %s\n", postAfterComment.CommentState) // 无评论
}


