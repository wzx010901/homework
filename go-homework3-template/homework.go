package homework03

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// -------------------------- 模型定义 --------------------------
// User 模型：用户（一对多关联Post）
type User struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`      // 主键，自增
	Username  string         `gorm:"size:50;not null;unique" json:"username"` // 用户名，唯一非空
	Email     string         `gorm:"size:100;unique" json:"email"`            // 邮箱，唯一
	Age       int            `gorm:"default:0" json:"age"`                    // 年龄，默认0
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`        // 自动创建时间
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`        // 自动更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                          // 软删除字段，索引

	// 一对多关联：一个用户有多个帖子
	Posts []Post `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"posts"`
	// constraint:OnDelete:CASCADE 表示删除用户时，级联删除其所有帖子
}

// Post 模型：文章（属于User，一对多关联Comment）
type Post struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"` // 主键，自增
	Title     string         `gorm:"size:200;not null" json:"title"`     // 文章标题，非空
	Content   string         `gorm:"type:text;not null" json:"content"`  // 文章内容，文本类型
	UserID    uint           `gorm:"not null;index" json:"user_id"`      // 关联用户的外键，非空，索引
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`   // 自动创建时间
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`   // 自动更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                     // 软删除字段，索引

	// 反向关联：属于某个用户（可选，用于通过Post查询所属User）
	User User `gorm:"foreignKey:UserID" json:"user"`

	// 一对多关联：一篇文章有多个评论
	Comments []Comment `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE" json:"comments"`
}

// Comment 模型：评论（属于Post）
type Comment struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"` // 主键，自增
	Content   string         `gorm:"size:500;not null" json:"content"`   // 评论内容，非空
	PostID    uint           `gorm:"not null;index" json:"post_id"`      // 关联文章的外键，非空，索引
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`   // 自动创建时间
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`   // 自动更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                     // 软删除字段，索引

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
