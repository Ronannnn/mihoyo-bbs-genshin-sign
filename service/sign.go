package service

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"mihoyo-bbs-genshin-sign/config"
	"mihoyo-bbs-genshin-sign/model"
	"mihoyo-bbs-genshin-sign/util"
	"net/http"
	"strings"
	"time"
)

func GetSignInfo(uid, cookie string) (signInfo *model.SignInfo, err error) {
	signInfo = &model.SignInfo{}
	var req *http.Request
	if req, err = http.NewRequest("GET", config.SignBaseUrl+config.SignAwardInfoUri, nil); err != nil {
		return
	}
	util.AddUrlQueryParametersFromStruct(req, model.SignUrlParam{
		ActId:  config.ActId,
		Uid:    uid,
		Region: getRegionFromUid(uid),
	})
	util.AddHeadersFromMap(req, map[string]string{
		"Cookie":   cookie,
		"SignHost": config.SignHost,
	})

	// request and parse response
	var resp *http.Response
	if resp, err = http.DefaultClient.Do(req); err != nil {
		return
	}
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	var respData = model.Response{Data: signInfo}
	if err = json.Unmarshal(body, &respData); err != nil {
		return
	}
	if respData.Retcode != 0 {
		return nil, fmt.Errorf("retcode: %d, message: %s", respData.Retcode, respData.Message)
	}
	return
}

func GetSignAwardList() (signAwardList *model.SignAwardList, err error) {
	signAwardList = &model.SignAwardList{}
	var req *http.Request
	if req, err = http.NewRequest("GET", config.SignBaseUrl+config.SignAwardHomeUri, nil); err != nil {
		return
	}
	util.AddUrlQueryParametersFromStruct(req, model.SignAwardsInfoReqParam{
		ActId: config.ActId,
	})

	// request and parse response
	var resp *http.Response
	if resp, err = http.DefaultClient.Do(req); err != nil {
		return
	}
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	var respData = model.Response{Data: signAwardList}
	if err = json.Unmarshal(body, &respData); err != nil {
		return
	}
	if respData.Retcode != 0 {
		return nil, fmt.Errorf("retcode: %d, message: %s", respData.Retcode, respData.Message)
	}
	return
}

func Sign(uid, cookie string) (err error) {
	var req *http.Request
	if req, err = http.NewRequest("POST", config.SignBaseUrl+config.SignAwardSignUri, nil); err != nil {
		return
	}

	util.AddUrlQueryParametersFromStruct(req, model.SignUrlParam{
		ActId:  config.ActId,
		Uid:    uid,
		Region: getRegionFromUid(uid),
	})
	util.AddHeadersFromMap(req, map[string]string{
		"Cookie":            cookie,
		"SignHost":          config.SignHost,
		"x-rpc-client_type": config.XRpcClientType,
		"x-rpc-app_version": config.XRpcClientVersion,
		"x-rpc-device_id":   strings.Replace(uuid.NewString(), "-", "", -1),
		"DS":                getDs(),
	})

	// request and parse response
	var resp *http.Response
	if resp, err = http.DefaultClient.Do(req); err != nil {
		return
	}
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	var respData = model.Response{}
	if err = json.Unmarshal(body, &respData); err != nil {
		return
	}
	if respData.Retcode != 0 {
		return fmt.Errorf("retcode: %d, message: %s", respData.Retcode, respData.Message)
	}
	return
}

func getDs() string {
	t := time.Now().Unix()
	println(t)
	r := util.GetRandString(6)
	hash := md5.Sum([]byte(fmt.Sprintf("salt=%s&t=%d&r=%s", config.DsSalt, t, r)))
	c := hex.EncodeToString(hash[:])
	return fmt.Sprintf("%d,%s,%s", t, r, c)
}

// getRegionFromUid get region according to the format of uid
func getRegionFromUid(uid string) string {
	if uid[0] == '5' {
		return config.RegionCnQd
	} else {
		return config.RegionCnGf
	}
}
