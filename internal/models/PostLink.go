package models

import (
	"gorm.io/gorm"
)

type PostLink struct {
	ID           uint `gorm:"primaryKey;autoIncrement"` // 显式声明自增
	CreatedAt    int64
	UpdatedAt    int64
	DeletedAt    int64
	Date         string `gorm:"index"`
	Title        string
	Summary      string
	PostLink     string
	CommentsLink string
	HotComments  string
}

// CreatePostLink 创建新的文章记录
func CreatePostLink(db *gorm.DB, post *PostLink) error {
	return db.Create(post).Error
}

// GetPostLinkByID 通过ID获取文章
func GetPostLinkByID(db *gorm.DB, id uint) (*PostLink, error) {
	var post PostLink
	result := db.First(&post, id)
	return &post, result.Error
}

// GetPostLinksByDate 获取特定日期的所有文章
func GetPostLinksByDate(db *gorm.DB, date string) ([]PostLink, error) {
	var posts []PostLink
	result := db.Where("date = ?", date).Find(&posts)
	return posts, result.Error
}

// GetRecentPostLinks 获取最近的n篇文章
func GetRecentPostLinks(db *gorm.DB, limit int) ([]PostLink, error) {
	var posts []PostLink
	result := db.Order("created_at desc").Limit(limit).Find(&posts)
	return posts, result.Error
}

// UpdatePostLink 更新文章信息
func UpdatePostLink(db *gorm.DB, post *PostLink) error {
	return db.Save(post).Error
}

// DeletePostLink 删除文章
func DeletePostLink(db *gorm.DB, id uint) error {
	return db.Delete(&PostLink{}, id).Error
}

// FindOrCreateByLink 查找或创建文章记录（根据链接判断唯一性）
func FindOrCreateByLink(db *gorm.DB, post *PostLink) (*PostLink, bool, error) {
	var existingPost PostLink
	result := db.Where("link = ?", post.PostLink).First(&existingPost)

	// 如果找不到记录
	if result.Error == gorm.ErrRecordNotFound {
		// 创建新记录
		if err := db.Create(post).Error; err != nil {
			return nil, false, err
		}
		return post, true, nil // 返回创建的记录和true表示是新创建的
	} else if result.Error != nil {
		return nil, false, result.Error // 其他错误
	}

	// 返回已存在的记录
	return &existingPost, false, nil
}
