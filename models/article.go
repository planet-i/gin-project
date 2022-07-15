package models

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"` // 外键
	Tag   Tag `json:"tag"`

	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Content   string `json:"content"`
	CreatedBy string `json:"modified_by"`
	State     int    `json:"state"`
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}
	return false
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

func GetAritcles(pageNum int, pageSize int, maps interface{}) (Articles []Article) {
	// Preload就是一个预加载器，它会执行两条 SQL，
	// SELECT * FROM blog_articles
	// SELECT * FROM blog_tag WHERE id IN (1,2,3,4)
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&Articles)

	return
}

func GetAritcle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag) // 通过Related关联查询

	return
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})

	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}

// 使用Unscoped硬删除文章数据
func CleanAllArticle() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{})

	return true
}
