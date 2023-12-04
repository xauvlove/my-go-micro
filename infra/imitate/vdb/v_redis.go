package vdb

import (
	"fmt"
	"time"
)

type RDBValue struct {
	// 值类型
	V interface{}
	// 设置时间
	t time.Time
	// 过期时间
	exp int64
}

var rdb = make(map[string]*RDBValue)

func SetString(key string, value string) {
	rdbValue := RDBValue{V: value, t: time.Now(), exp: -1}
	rdb[key] = &rdbValue
}

func Set(key string, value interface{}) {
	rdbValue := RDBValue{V: value, t: time.Now(), exp: -1}
	rdb[key] = &rdbValue
}

func SetExString(key string, value string, expt int64) {
	rdbValue := RDBValue{V: value, t: time.Now(), exp: expt}
	rdb[key] = &rdbValue
}

func SetEx(key string, value interface{}, expt int64) {
	rdbValue := RDBValue{V: value, t: time.Now(), exp: expt}
	rdb[key] = &rdbValue
}

func Remove(key string) {
	delete(rdb, key)
}

func GetString(key string) string {
	v, ok := rdb[key]
	if ok {
		if v.t.UnixMilli()+v.exp*1000 < time.Now().UnixMilli() {
			Remove(key)
			return ""
		}
		return fmt.Sprintf("%v", v.V)
	}
	return ""
}

func Get(key string) interface{} {
	v, ok := rdb[key]
	if ok {
		if v.t.UnixMilli()+v.exp*1000 < time.Now().UnixMilli() {
			Remove(key)
			return nil
		}
		return v.V
	}
	return nil
}
