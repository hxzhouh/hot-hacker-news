package handlers

import (
	"hot-hacker-new/internal/database"
	"hot-hacker-new/internal/models"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewsHandler struct {
	Template *template.Template
}

func NewNewsHandler(tmpl *template.Template) *NewsHandler {
	return &NewsHandler{Template: tmpl}
}

func (h *NewsHandler) Index(c *gin.Context) {
	// 获取最近一天的新闻。
	_, posts, err := models.FindLastDatePost(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取新闻"})
		return
	}
	c.HTML(http.StatusOK, "news/index.html", gin.H{"posts": posts})
}

func (h *NewsHandler) Detail(c *gin.Context) {
	// id := c.Param("id")
	// post, err := models.GetPostByID(id) // 假设有一个根据ID获取新闻的函数
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "新闻未找到"})
	// 	return
	// }
	// c.HTML(http.StatusOK, "news/detail.html", gin.H{"post": post})
}
