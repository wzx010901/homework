package homework03

import (
	"fmt"
	"sync/atomic"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// -------------------------- 模型定义 --------------------------
// User 模型：用户（一对多关联Post）
type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`      // 主键，自增
	Username  string    `gorm:"size:50;not null;unique" json:"username"` // 用户名，唯一非空
	Email     string    `gorm:"size:100;unique" json:"email"`            // 邮箱，唯一
	Age       int       `gorm:"default:0" json:"age"`                    // 年龄，默认0
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`        // 自动创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`        // 自动更新时间

	// 一对多关联：一个用户有多个帖子
	Posts []Post `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"posts"`
	// constraint:OnDelete:CASCADE 表示删除用户时，级联删除其所有帖子

	PostCount int64 `gorm:"default:0"` // 文章数量统计字段
}

// Post 模型：文章（属于User，一对多关联Comment）
type Post struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"` // 主键，自增
	Title         string    `gorm:"size:200;not null" json:"title"`     // 文章标题，非空
	Content       string    `gorm:"type:text;not null" json:"content"`  // 文章内容，文本类型
	UserID        uint      `gorm:"not null;index" json:"user_id"`      // 关联用户的外键，非空，索引
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`   // 自动创建时间
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`   // 自动更新时间
	CommentStatus string    `gorm:"default:'有评论'"`                      // 评论状态：有评论/无评论
	// 反向关联：属于某个用户（可选，用于通过Post查询所属User）
	User User `gorm:"foreignKey:UserID" json:"user"`

	// 一对多关联：一篇文章有多个评论
	Comments []Comment `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE" json:"comments"`
}

// Comment 模型：评论（属于Post）
type Comment struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"` // 主键，自增
	Content      string    `gorm:"size:500;not null" json:"content"`   // 评论内容，非空
	PostID       uint      `gorm:"not null;index" json:"post_id"`      // 关联文章的外键，非空，索引
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`   // 自动创建时间
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`   // 自动更新时间
	CommentCount int64     `gorm:"-"`                                  // 非数据库字段，用于存储评论数
	// 反向关联：属于某篇文章（可选）
	Post Post `gorm:"foreignKey:PostID" json:"post"`
}

// 编写Go代码，使用Gorm创建这些模型对应的数据库表。
func CreateTables() {
	// -------------------------- 初始化Gorm连接 --------------------------
	// 连接SQLite数据库（文件名为gorm_demo.db，不存在则自动创建）
	db, err := gorm.Open(sqlite.Open("gorm_demo2.db"), &gorm.Config{
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
	// 创建用户
	user := User{
		Username: "zhangsan",
		Email:    "zhangsan@example.com",
		Age:      28,
	}
	db.Create(&user)
	fmt.Println("创建用户成功，ID：", user.ID)

	// 为用户创建帖子
	post := Post{
		Title:   "Gorm一对多关联示例",
		Content: "这是一篇演示Gorm关联关系的文章",
		UserID:  user.ID, // 关联用户ID
	}
	db.Create(&post)
	fmt.Println("创建帖子成功，ID：", post.ID)

	// 为帖子创建评论
	comments := []Comment{
		{Content: "这篇文章写得很好！", PostID: post.ID},
		{Content: "感谢分享，学到了！", PostID: post.ID},
	}
	db.Create(&comments)
	fmt.Println("创建评论成功，数量：", len(comments))

}

// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
func QueryUserPostsAndComments() {
	db, err := gorm.Open(sqlite.Open("gorm_demo.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 打印SQL日志，方便调试
	})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
	//  查询用户及其所有帖子（预加载关联）
	var queryUser User
	db.Preload("Posts.Comments").First(&queryUser, 1) // 预加载帖子+帖子下的评论
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

type PostCommentCount struct {
	PostID       uint  // 文章ID（对应comment表的post_id）
	CommentCount int64 // 评论数（对应COUNT(*)的别名）
}

// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
func QueryMostCommentedPost() {
	db, err := gorm.Open(sqlite.Open("gorm_demo.db"), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
	var post Post
	var comment PostCommentCount
	error2 := db.Model(&Comment{}).Select("post_id, COUNT(*) as comment_count").Group("post_id").Order("comment_count DESC").Limit(1).First(&comment)
	//打印sql
	fmt.Println(db.Session(&gorm.Session{}).Statement.SQL.String())
	if error2 != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Errorf("暂无文章数据")
		}
		fmt.Errorf("查询失败：%v", err)
	}
	err1 := db.Model(&Post{}).Where("id = ?", comment.PostID).First(&post)
	if err1 != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Errorf("暂无文章数据")
		}
		fmt.Errorf("查询失败：%v", err)
	}

	fmt.Printf("\n评论数量最多的文章：%s,数量是：%d", post.Title, comment.CommentCount)

}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段
func (p *Post) AfterCreate(tx *gorm.DB) error {
	// 校验用户ID是否有效
	if p.UserID == 0 {
		return fmt.Errorf("用户ID为空，无法更新文章数量")
	}
	// 原子操作（推荐，并发安全）
	var user User
	if err := tx.First(&user, p.UserID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("用户ID %d 不存在，无法更新文章数量", p.UserID)
		}
		return fmt.Errorf("查询用户失败：%v", err)
	}
	atomic.AddInt64(&user.PostCount, 1) // 原子递增
	err := tx.Model(&user).Update("post_count", user.PostCount).Error

	// 处理更新失败的情况
	if err != nil {
		return fmt.Errorf("更新用户文章数量失败：%v", err)
	}

	fmt.Printf("钩子执行成功：用户ID %d 的文章数量已更新为 %d\n", p.UserID, user.PostCount)
	return nil
}

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm_demo_hook.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 打印SQL日志，方便调试
	})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
	if err != nil {
		panic(fmt.Sprintf("数据库连接失败：%v", err))
	}

	// 自动迁移表结构（包含新增的PostCount字段）
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		panic(fmt.Sprintf("表迁移失败：%v", err))
	}
	return db
}

func Hook() {
	db := initDB()

	//先创建一个测试用户
	user := User{Username: "hook_test", Email: "hook@example.com"}
	if err := db.Create(&user).Error; err != nil {
		fmt.Printf("创建用户失败：%v\n", err)
		return
	}
	fmt.Printf("初始用户信息：ID=%d，用户名=%s，文章数量=%d\n", user.ID, user.Username, user.PostCount)

	//  创建文章（触发AfterCreate钩子）
	post1 := Post{Title: "钩子测试文章1", Content: "文章内容1", UserID: user.ID}
	post2 := Post{Title: "钩子测试文章2", Content: "文章内容2", UserID: user.ID}
	if err := db.Create(&post1).Error; err != nil {
		fmt.Printf("创建文章1失败：%v\n", err)
	}
	if err := db.Create(&post2).Error; err != nil {
		fmt.Printf("创建文章2失败：%v\n", err)
	}

	//  查询用户最新信息，验证文章数量是否更新
	var updatedUser User
	if err := db.First(&updatedUser, user.ID).Error; err != nil {
		fmt.Printf("查询用户失败：%v\n", err)
		return
	}
	fmt.Println("\n===== 最终用户信息 =====")
	fmt.Printf("用户ID：%d\n", updatedUser.ID)
	fmt.Printf("用户名：%s\n", updatedUser.Username)
	fmt.Printf("文章数量：%d\n", updatedUser.PostCount) // 预期值：2
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	// 校验文章ID是否有效
	if c.PostID == 0 {
		return fmt.Errorf("文章ID为空，无法检查评论数量")
	}

	// 查询对应文章（确保文章存在）
	var post Post
	if err := tx.First(&post, c.PostID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("文章ID %d 不存在，无需更新评论状态", c.PostID)
		}
		return fmt.Errorf("查询文章失败：%v", err)
	}

	// 统计该文章剩余的评论数量
	var commentCount int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount).Error; err != nil {
		return fmt.Errorf("统计文章评论数失败：%v", err)
	}

	// 根据评论数更新文章的评论状态
	var newStatus string
	if commentCount == 0 {
		newStatus = "无评论"
	} else {
		newStatus = "有评论" // 仍有评论，恢复/保持“有评论”状态
	}

	// 更新文章状态
	if err := tx.Model(&post).Update("comment_status", newStatus).Error; err != nil {
		return fmt.Errorf("更新文章评论状态失败：%v", err)
	}

	fmt.Printf("钩子执行成功：文章ID %d 的评论数为 %d，评论状态已更新为「%s」\n", c.PostID, commentCount, newStatus)
	return nil
}
func initDB2() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm_demo_hook2.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 打印SQL日志，方便调试
	})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
	if err != nil {
		panic(fmt.Sprintf("数据库连接失败：%v", err))
	}

	// 自动迁移表结构（包含新增的PostCount字段）
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		panic(fmt.Sprintf("表迁移失败：%v", err))
	}
	return db
}
func Hook2() {
	db := initDB2()

	// 初始化测试数据：创建用户→创建文章→创建评论
	// 创建用户
	user := User{Username: "comment_hook_test", Email: "comment_hook@example.com"}
	if err := db.Create(&user).Error; err != nil {
		fmt.Printf("创建用户失败：%v\n", err)
		return
	}

	// 创建文章（初始评论状态为“有评论”）
	post := Post{Title: "评论状态测试文章", Content: "测试内容", UserID: user.ID}
	if err := db.Create(&post).Error; err != nil {
		fmt.Printf("创建文章失败：%v\n", err)
		return
	}
	fmt.Printf("初始文章信息：ID=%d，标题=%s，评论状态=%s\n", post.ID, post.Title, post.CommentStatus)

	//  创建2条评论
	comments := []Comment{
		{Content: "测试评论1", PostID: post.ID},
		{Content: "测试评论2", PostID: post.ID},
	}

	if err := db.Create(&comments).Error; err != nil {
		fmt.Printf("创建评论失败：%v\n", err)
		return
	}
	fmt.Println("已创建2条测试评论")

	// 删除第一条评论（触发AfterDelete钩子，此时剩余1条评论，状态仍为“有评论”）

	if err := db.Delete(&comments[0]).Error; err != nil {
		fmt.Printf("删除评论 1 失败：%v\n", err)
		return
	}

	// 删除第二条评论（触发AfterDelete钩子，剩余0条评论，状态更新为“无评论”）

	if err := db.Delete(&comments[1]).Error; err != nil {
		fmt.Printf("删除评论 1 失败：%v\n", err)
		return
	}

	//  查询文章最新状态，验证评论状态是否更新
	var updatedPost Post
	if err := db.First(&updatedPost, post.ID).Error; err != nil {
		fmt.Printf("查询文章失败：%v\n", err)
		return
	}
	fmt.Println("\n===== 最终文章信息 =====")
	fmt.Printf("文章ID：%d\n", updatedPost.ID)
	fmt.Printf("标题：%s\n", updatedPost.Title)
	fmt.Printf("评论状态：%s\n", updatedPost.CommentStatus) // 预期值：无评论
}
