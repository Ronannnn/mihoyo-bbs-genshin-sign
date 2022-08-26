package model

import "gorm.io/gorm"

type SignItem struct {
	Base
	Uid    string `json:"uid"`
	Cookie string `json:"cookie"`
	Email  string `json:"email"`
}

func (signItem *SignItem) Create(db *gorm.DB) error {
	return db.Create(signItem).Error
}

func (signItem *SignItem) Update(db *gorm.DB) error {
	return db.Updates(signItem).Error
}

func FindSignItemById(db *gorm.DB, id int) (signItem SignItem, err error) {
	err = db.Find(&signItem, "id = ?", id).Error
	return
}

func FindAllSignItems(db *gorm.DB) (signItems []SignItem, err error) {
	err = db.Find(&signItems).Error
	return
}

func DeleteSignItemById(db *gorm.DB, signItemId int) error {
	return db.Delete(SignItem{}, "id = ?", signItemId).Error
}
