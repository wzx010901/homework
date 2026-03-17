package handlers

import (
	"blog-system/config"
	"blog-system/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var req struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		config.Warn("创建文章参数错误:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误：" + err.Error(),
		})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未授权",
		})
		return
	}

	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID.(uint),
	}

	if err := models.DB.Create(&post).Error; err != nil {
		config.Error("创建文章失败:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建文章失败",
		})
		return
	}

	config.Info("文章创建成功:", post.Title)
	c.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": "文章创建成功",
		"data":    post,
	})
}

func GetPosts(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	var total int64
	if err := models.DB.Model(&models.Post{}).Count(&total).Error; err != nil {
		config.Error("获取文章总数失败:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取文章列表失败",
		})
		return
	}

	var posts []models.Post
	offset := (page - 1) * pageSize
	if err := models.DB.Preload("User").Order("created_at DESC").Limit(pageSize).Offset(offset).Find(&posts).Error; err != nil {
		config.Error("获取文章列表失败:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取文章列表失败",
		})
		return
	}
	type PostReponse struct {
		ID        uint              `gorm:"primaryKey;autoIncrement" json:"id"`
		Title     string            `gorm:"size:200;not null" json:"title"`
		Content   string            `gorm:"type:text;not null" json:"content"`
		UserID    uint              `gorm:"not null;index" json:"user_id"`
		CreatedAt models.CustomTime `gorm:"autoCreateTime" json:"created_at"`
		UpdatedAt models.CustomTime `gorm:"autoUpdateTime" json:"updated_at"`
	}
	var postsReponse []PostReponse
	for _, post := range posts {
		postsReponse = append(postsReponse, PostReponse{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			UserID:    post.UserID,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})

	}

	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"message":  "获取文章列表成功",
		"data":     postsReponse,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

func GetPost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的文章 ID",
		})
		return
	}

	var post models.Post
	if err := models.DB.Preload("User").Preload("Comments.User").First(&post, id).Error; err != nil {
		config.Warn("文章不存在:", id)
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "文章不存在",
		})
		return
	}
	type PostReponse struct {
		ID        uint              `gorm:"primaryKey;autoIncrement" json:"id"`
		Title     string            `gorm:"size:200;not null" json:"title"`
		Content   string            `gorm:"type:text;not null" json:"content"`
		UserID    uint              `gorm:"not null;index" json:"user_id"`
		CreatedAt models.CustomTime `gorm:"autoCreateTime" json:"created_at"`
		UpdatedAt models.CustomTime `gorm:"autoUpdateTime" json:"updated_at"`
	}
	postsReponse := PostReponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取文章成功",
		"data":    postsReponse,
	})
}

func UpdatePost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的文章 ID",
		})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未授权",
		})
		return
	}

	var post models.Post
	if err := models.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "文章不存在",
		})
		return
	}

	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "无权限修改他人的文章",
		})
		return
	}

	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
		})
		return
	}

	if req.Title != "" {
		post.Title = req.Title
	}
	if req.Content != "" {
		post.Content = req.Content
	}

	if err := models.DB.Save(&post).Error; err != nil {
		config.Error("更新文章失败:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新文章失败",
		})
		return
	}

	config.Info("文章更新成功:", post.Title)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "文章更新成功",
		"data":    post,
	})
}

func DeletePost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的文章 ID",
		})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未授权",
		})
		return
	}

	var post models.Post
	if err := models.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "文章不存在",
		})
		return
	}

	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "无权限删除他人的文章",
		})
		return
	}

	if err := models.DB.Delete(&post).Error; err != nil {
		config.Error("删除文章失败:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除文章失败",
		})
		return
	}

	config.Info("文章删除成功:", post.Title)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "文章删除成功",
	})
}
