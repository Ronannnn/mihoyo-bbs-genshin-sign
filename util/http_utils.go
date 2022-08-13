package util

import (
	"fmt"
	"mihoyo-bbs-genshin-sign/config"
	"net/http"
	"reflect"
)

func AddUrlQueryParametersFromStruct(req *http.Request, data interface{}) {
	reqQuery := req.URL.Query()
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(config.HttpQueryTagName)
		reqQuery.Add(tag, fmt.Sprintf("%v", v.Field(i).Interface()))
	}
	req.URL.RawQuery = reqQuery.Encode()
}

func AddUrlQueryParametersFromMap(req *http.Request, data map[string]string) {
	reqQuery := req.URL.Query()
	for k, v := range data {
		reqQuery.Add(k, v)
	}
}

func AddHeadersFromStruct(req *http.Request, data interface{}) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(config.HttpQueryTagName)
		req.Header.Set(tag, fmt.Sprintf("%v", v.Field(i).Interface()))
	}
}

func AddHeadersFromMap(req *http.Request, data map[string]string) {
	for k, v := range data {
		req.Header.Set(k, v)
	}
}
