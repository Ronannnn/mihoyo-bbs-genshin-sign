package service_test

import (
	"github.com/stretchr/testify/assert"
	"mihoyo-bbs-genshin-sign/service"
	"testing"
)

const (
	uid    = "your uid"
	cookie = "your cookie"
)

func TestGetSignInfo(t *testing.T) {
	_, err := service.GetSignInfo(uid, cookie)
	assert.NoError(t, err)
}

func TestGetSignAwardList(t *testing.T) {
	ret, err := service.GetSignAwardList()
	assert.NoError(t, err)
	assert.True(t, ret.Month >= 1 && ret.Month <= 12)
}

func TestSign(t *testing.T) {
	err := service.Sign(uid, cookie)
	assert.NoError(t, err)
}
