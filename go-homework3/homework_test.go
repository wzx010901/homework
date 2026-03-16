package homework03

import (
	"fmt"
	"sync"
	"testing"

	"gorm.io/driver/sqlite" // 使用SQLite演示（无需额外配置，文件存储）
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	failedQuestions []string
	totalQuestions  int
	mu              sync.Mutex
)

func recordResult(t *testing.T, name string) {
	mu.Lock()
	defer mu.Unlock()
	totalQuestions++
	if t.Failed() {
		failedQuestions = append(failedQuestions, name)
	}
}

func TestOne(t *testing.T) {
	// -------------------------- 初始化Gorm连接 --------------------------
	// 连接SQLite数据库（文件名为gorm_demo.db，不存在则自动创建）
	db, err := gorm.Open(sqlite.Open("gorm_demo.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 打印SQL日志，方便调试
	})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}

	// -------------------------- 自动迁移表结构 --------------------------
	// 按模型创建/更新表，会自动处理外键和关联关系
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		panic("表结构迁移失败：" + err.Error())
	}
	fmt.Println("表结构迁移成功！")

	// -------------------------- 示例：创建关联数据 --------------------------
	// 1. 创建用户
	user := User{
		Username: "zhangsan",
		Email:    "zhangsan@example.com",
		Age:      28,
	}
	db.Create(&user)
	fmt.Println("创建用户成功，ID：", user.ID)

	// 2. 为用户创建帖子
	post := Post{
		Title:   "Gorm一对多关联示例",
		Content: "这是一篇演示Gorm关联关系的文章",
		UserID:  user.ID, // 关联用户ID
	}
	db.Create(&post)
	fmt.Println("创建帖子成功，ID：", post.ID)

	// 3. 为帖子创建评论
	comments := []Comment{
		{Content: "这篇文章写得很好！", PostID: post.ID},
		{Content: "感谢分享，学到了！", PostID: post.ID},
	}
	db.Create(&comments)
	fmt.Println("创建评论成功，数量：", len(comments))

	// -------------------------- 示例：查询关联数据 --------------------------
	// 1. 查询用户及其所有帖子（预加载关联）
	var queryUser User
	db.Preload("Posts.Comments").First(&queryUser, user.ID) // 预加载帖子+帖子下的评论
	fmt.Println("\n查询到用户：", queryUser.Username)
	fmt.Println("用户的帖子数：", len(queryUser.Posts))
	for _, p := range queryUser.Posts {
		fmt.Println("  帖子标题：", p.Title)
		fmt.Println("  帖子评论数：", len(p.Comments))
		for _, c := range p.Comments {
			fmt.Println("    评论内容：", c.Content)
		}
	}
}

func TestCreateTables(t *testing.T) {
	CreateTables()
}

func TestQueryUserPostsAndComments(t *testing.T) {
	QueryUserPostsAndComments()
}

func TestQueryMostCommentedPost(t *testing.T) {
	QueryMostCommentedPost()
}

func TestHook(t *testing.T) {
	Hook()
}

func TestHook2(t *testing.T) {
	Hook2()
}
