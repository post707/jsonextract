package jsonextract

import (
	"encoding/json"
	"errors"
	"strconv"
)

type JS struct {
	Data interface{}
}

func MyJson(js string) *JS {
	j := new(JS)
	var i interface{}
	err := json.Unmarshal([]byte(js), &i)
	if err != nil {
		return j
	}
	j.Data = i
	return j
}

// func MyJstomap(js string) map[string]interface{} {
// 	var i map[string]interface{}
// 	err := json.Unmarshal([]byte(js), &i)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return i
// }

func (myjson *JS) Getkey(key string) *JS {
	if v, ok := (myjson.GetData())[key]; ok {
		myjson.Data = v
		return myjson
	}
	myjson.Data = nil
	return myjson
}

//获取数组中低num的键值
func (myjson *JS) GetArrkey(key string, i int) *JS {
	num := i - 1
	if i > len((myjson.Data).([]interface{})) {
		myjson.Data = errors.New("index out of range list").Error()
		return myjson
	}
	if m, ok := (myjson.Data).([]interface{}); ok {
		v := m[num].(map[string]interface{})
		if h, ok := v[key]; ok {
			myjson.Data = h
			return myjson
		}

	}
	myjson.Data = nil
	return myjson
}

func (myjson *JS) GetData() map[string]interface{} {
	if m, ok := (myjson.Data).(map[string]interface{}); ok {
		return m
	}
	return nil
}
func (myjson *JS) ToString() string {
	if m, ok := myjson.Data.(string); ok {
		return m
	}
	if m, ok := myjson.Data.(float64); ok {
		return strconv.FormatFloat(m, 'f', -1, 64)
	}
	return ""
}

// a := strconv.FormatFloat(10.010, 'f', -1, 64)
// 输出：10.01

func (myjson *JS) GetArrylen() int {
	return len((myjson.Data).([]interface{}))
}
