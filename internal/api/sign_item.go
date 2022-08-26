package api

import (
	"mihoyo-bbs-genshin-sign/internal/config"
	"mihoyo-bbs-genshin-sign/internal/model"
	"mihoyo-bbs-genshin-sign/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllSignItems(c *gin.Context) {
	if signItems, err := model.FindAllSignItems(db); err != nil {
		FailWithErr(c, err)
	} else {
		OkWithData(c, signItems)
	}
}

func CreateSignItem(c *gin.Context) {
	var err error
	var signItem model.SignItem
	if err = c.BindJSON(&signItem); err != nil {
		FailWithErr(c, err)
		return
	}
	if signItem.Id != 0 {
		FailWithErr(c, model.ErrCreatedId)
		return
	}
	if err = signItem.Create(db); err != nil {
		FailWithErr(c, err)
		return
	}
	OkWithData(c, signItem)
}

func UpdateSignItem(c *gin.Context) {
	var err error
	var tmpSignItem model.SignItem
	var oprId = c.GetInt(config.CtxKeyUserId)
	if err = c.BindJSON(&tmpSignItem); err != nil {
		FailWithErr(c, err)
		return
	}
	tmpSignItem.UpdatedBy = oprId
	var updatedSignItem model.SignItem
	if updatedSignItem, err = service.UpdateSignItem(db, tmpSignItem); err != nil {
		FailWithErr(c, err)
		return
	}
	OkWithData(c, updatedSignItem)
}

func DeleteSignItemById(c *gin.Context) {
	var signItemId int
	var err error
	if signItemId, err = strconv.Atoi(c.Param("id")); err != nil {
		FailWithErr(c, err)
		return
	}
	if err = service.DeleteSignItem(db, signItemId); err != nil {
		FailWithErr(c, err)
		return
	}
	Ok(c)
}
