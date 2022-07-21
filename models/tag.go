package models

import "gorm.io/gorm"

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// ExistTagByName 检查是否存在同名标签
func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ? AND deleted_on = ? ", name, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

// ExistTagByID 根据ID判断标签是否存在
func ExistTagByID(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, nil
	}

	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetTags 根据分页和条件获取标签列表
func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	var (
		tags []Tag
		err  error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&tags).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&tags).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

// GetTagTotal 根据条件获取标签总数
func GetTagTotal(maps interface{}) (int, error) {
	var count int
	err := db.Model(&Tag{}).Where(maps).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// AddTag 添加标签
func AddTag(name string, state int, createdBy string) error {
	tag := &Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}

	if err := db.Create(&tag).Error; err != nil {
		return err
	}

	return nil
}

// DeleteTag 删除标签
func DeleteTag(id int) error {
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}

	return nil
}

// EditTag 编辑单个标签
func EditTag(id int, data interface{}) error {
	if err := db.Model(&Tag{}).Where("id = ? AND deleted_on = ? ", id, 0).Update(data).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllTag 清理被删除的标签数据
// 使用Unscoped硬删除标签数据
func CleanAllTag() (bool, error) {
	if err := db.Unscoped().Where("delete_on != ?", 0).Delete(&Tag{}).Error; err != nil {
		return false, err
	}

	return true, nil
}
