# Go Homework 01 Template

## 使用说明

1. 编辑 `homework.go`，完成各个函数的实现。
2. 在本地运行 `go test -v` 验证代码。
3. 提交代码到你的分支。GitHub Actions 会自动运行测试。

## 题目列表

1. CreateTables (编写Go代码，使用Gorm创建这些模型对应的数据库表。)
2. QueryUserPostsAndComments (编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。)
3. QueryMostCommentedPost (编写Go代码，使用Gorm查询评论数量最多的文章信息。)
4. Hook (为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段)
5. Hook2 (为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论")

