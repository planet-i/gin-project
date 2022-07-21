package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"` // 外键
	Tag   Tag `json:"tag"`                 // 标签

	Title         string `json:"title"`           // 标题
	Desc          string `json:"desc"`            // 简述
	Content       string `json:"content"`         // 内容
	CoverImageUrl string `json:"cover_image_url"` // 封面
	CreatedBy     string `json:"created_by"`      // 创建者
	ModifiedBy    string `json:"modified_by"`     // 修改者
	State         int    `json:"state"`           // 状态
}

// ExistArticleByID 根据ID检查文章是否存在
func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetArticleTotal 根据条件获取文章总数
func GetArticleTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Article{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetArticles 根据页码获取文章列表
func GetArticles(pageNum int, pageSize int, maps interface{}) ([]*Article, error) {
	var articles []*Article
	err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articles, nil
}

// GetArticle 获取指定文章
func GetArticle(id int) (*Article, error) {
	var article *Article
	// 通过Related关联查询
	err := db.Where("id = ? AND delete_on = ?", id, 0).First(article).Related(&article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return article, nil
}

// EditArticle 修改指定文章
func EditArticle(id int, data interface{}) error {
	if err := db.Model(&Article{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// AddArticle 添加文章
func AddArticle(data map[string]interface{}) error {
	article := Article{
		TagID:         data["tag_id"].(int),
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["content"].(string),
		CreatedBy:     data["created_by"].(string),
		State:         data["state"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	}
	if err := db.Create(&article).Error; err != nil {
		return err
	}

	return nil
}

// DeleteArticle 删除指定文章
func DeleteArticle(id int) error {
	if err := db.Where("id = ?", id).Delete(Article{}).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllArticle 清理被软删除的数据
func CleanAllArticle() error {
	// 使用Unscoped硬删除
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{}).Error; err != nil {
		return err
	}

	return nil
}
