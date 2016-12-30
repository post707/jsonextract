package jsonextract

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type JS struct {
	Data interface{}
}

func Json1(js string) map[string]interface{} {
	var i map[string]interface{}
	err := json.Unmarshal([]byte(js), &i)
	if err != nil {
		fmt.Println(err)
	}
	return i
}

func Json(js string) *JS {
	j := new(JS)
	var i interface{}
	err := json.Unmarshal([]byte(js), &i)
	if err != nil {
		return j
	}
	j.Data = i
	return j
}

func (dejson *JS) Getkey(key string) *JS {
	if v, ok := (dejson.GetData())[key]; ok {
		dejson.Data = v
		return dejson
	}
	dejson.Data = nil
	return dejson
}

func (dejson *JS) GetData() map[string]interface{} {
	if m, ok := (dejson.Data).(map[string]interface{}); ok {
		return m
	}
	return nil
}

func (dejson *JS) ToString() string {
	if m, ok := dejson.Data.(string); ok {
		return m
	}
	if m, ok := dejson.Data.(float64); ok {
		return strconv.FormatFloat(m, 'f', -1, 64)
	}
	return ""
}

func (dejson *JS) GetArrylen() int {
	return len((dejson.Data).([]interface{}))
}
