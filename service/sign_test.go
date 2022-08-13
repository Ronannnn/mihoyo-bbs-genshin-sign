package service_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"mihoyo-bbs-genshin-sign/service"
	"testing"
)

func TestGetSignInfo(t *testing.T) {
	_, err := service.GetSignInfo("your uid", "your cookie")
	assert.NoError(t, err)
}

func TestGetSignAwardList(t *testing.T) {
	ret, err := service.GetSignAwardList()
	assert.NoError(t, err)
	assert.True(t, ret.Month >= 1 && ret.Month <= 12)
}

func TestSign(t *testing.T) {
	err := service.Sign("your uid", "your cookie")
	fmt.Println(err)
}
