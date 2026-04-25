package handlers

import (
	"blog-system/config"
	"blog-system/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var req struct {
		PostID  uint   `json:"postId" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		config.Warn("创建评论参数错误:", err)
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

	var post models.Post
	if err := models.DB.First(&post, req.PostID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "文章不存在",
		})
		return
	}

	comment := models.Comment{
		Content: req.Content,
		UserID:  userID.(uint),
		PostID:  req.PostID,
	}

	if err := models.DB.Create(&comment).Error; err != nil {
		config.Error("创建评论失败:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建评论失败",
		})
		return
	}

	config.Info("评论创建成功")
	c.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": "评论创建成功",
		"data":    comment,
	})
}

func GetCommentsByPost(c *gin.Context) {
	postIDStr := c.Param("postID")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的文章 ID",
		})
		return
	}

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
	if err := models.DB.Model(&models.Comment{}).Where("post_id = ?", postID).Count(&total).Error; err != nil {
		config.Error("获取评论总数失败:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取评论列表失败",
		})
		return
	}

	var comments []models.Comment
	offset := (page - 1) * pageSize
	if err := models.DB.Where("post_id = ?", postID).Preload("User").Order("created_at ASC").Limit(pageSize).Offset(offset).Find(&comments).Error; err != nil {
		config.Error("获取评论列表失败:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取评论列表失败",
		})
		return
	}
	type CommentReponse struct {
		ID        uint              `gorm:"primaryKey;autoIncrement" json:"id"`
		Content   string            `gorm:"size:500;not null" json:"content"`
		UserID    uint              `gorm:"not null;index" json:"user_id"`
		PostID    uint              `gorm:"not null;index" json:"post_id"`
		CreatedAt models.CustomTime `gorm:"autoCreateTime" json:"created_at"`
		UpdatedAt models.CustomTime `gorm:"autoUpdateTime" json:"updated_at"`
	}
	var commentsReponse []CommentReponse
	for _, comment := range comments {
		commentsReponse = append(commentsReponse, CommentReponse{
			ID:        comment.ID,
			Content:   comment.Content,
			UserID:    comment.UserID,
			PostID:    comment.PostID,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"message":  "获取评论列表成功",
		"data":     commentsReponse,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}
