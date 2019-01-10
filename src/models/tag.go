package models

import (
	"fmt"
)

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func ExistTagByID(id int) bool {
	var tag Tag
	err := db.Select("id").Where("id = ?", id).First(&tag).Error
	if err != nil {
		return false
	} else {
		if tag.ID > 0 {
			return true
		} else {
			return false
		}
	}
}

func AddTag(name string, state int, createdBy string) bool {
	tag := &Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}
	err := db.Create(tag).Error
	if err != nil {
		fmt.Printf("插入tag出错，%v", err)
		return false
	}

	return db.NewRecord(tag)
}

func DeleteTag(id int) bool {
	err := db.Where("id = ?", id).Delete(&Tag{}).Error
	if err != nil {
		return false
	} else {
		return true
	}
}

func EditTag(id int, data interface{}) bool {
	err := db.Model(&Tag{}).Where("id = ?", id).Update(data).Error
	if err != nil {
		return false
	} else {
		return true
	}
}
