package service

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"mihoyo-bbs-genshin-sign/model"
	"time"
)

func UpdateSignItem(db *gorm.DB, after model.SignItem) (updatedSignItem model.SignItem, err error) {
	if after.Id == 0 {
		return updatedSignItem, model.ErrUpdatedId
	}
	if err = db.Transaction(func(tx *gorm.DB) (err error) {
		var before model.SignItem
		// if not lock the row, another transaction may get this row at the same time
		// and these two transactions can both succeed to modify this row
		if err = tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&before, "id = ?", after.Id).
			Error; err != nil {
			return
		}
		if !after.UpdatedAt.Round(time.Second).Equal(before.UpdatedAt) {
			return model.ErrModified
		}
		after.UpdatedAt = time.Now()
		if err = after.Update(tx); err != nil {
			return
		}
		return
	}); err != nil {
		return updatedSignItem, err
	}
	return model.FindSignItemById(db, after.Id)
}

func DeleteSignItem(db *gorm.DB, userId int) (err error) {
	return db.Transaction(func(tx *gorm.DB) (err error) {
		// if not lock the row, another transaction may get this row at the same time
		// and these two transactions can both succeed to modify this row
		if err = tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&model.SignItem{}, "id = ?", userId).
			Error; err != nil {
			return
		}
		// delete user
		if err = model.DeleteSignItemById(tx, userId); err != nil {
			return
		}
		return
	})
}
