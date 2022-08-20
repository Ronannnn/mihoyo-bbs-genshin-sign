package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response According to https://pro.ant.design/zh-CN/docs/request
type Response struct {
	Success      bool        `json:"success"`
	Data         interface{} `json:"data"`
	ErrorCode    string      `json:"errorCode"`
	ErrorMessage string      `json:"errorMessage"`
	ShowType     ShowType    `json:"showType"`
}

const (
	FAIL    = false
	SUCCESS = true
)

type ShowType int

//goland:noinspection GoUnusedConst
const (
	Silent       ShowType = 0
	WarnMessage  ShowType = 1
	ErrorMessage ShowType = 2
	Notification ShowType = 4
	Redirect     ShowType = 9
)

const (
	NormalError = "0"
	TokenError  = "1"
)

func Result(success bool, data interface{}, errMsg string, showType ShowType, c *gin.Context) {
	log.Debugf("success: %v, data: %v, errMsg: %s", success, data, errMsg)
	c.JSON(http.StatusOK, Response{
		success,
		data,
		NormalError,
		errMsg,
		showType,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "", Silent, c)
}

func OkWithData(c *gin.Context, data interface{}) {
	Result(SUCCESS, data, "", Silent, c)
}

func FailWithMsg(c *gin.Context, msg string) {
	Result(FAIL, map[string]interface{}{}, msg, ErrorMessage, c)
}

func FailWithErr(c *gin.Context, err error) {
	Result(FAIL, map[string]interface{}{}, err.Error(), ErrorMessage, c)
}

func FailWithData(c *gin.Context, data interface{}) {
	Result(FAIL, data, "", ErrorMessage, c)
}
