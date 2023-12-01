package vdb

import "time"

type RDBValue struct {
	// 值类型
	V string
	// 设置时间
	t time.Time
	// 过期时间
	exp int64
}

var rdb = make(map[string]*RDBValue)

func Set(key string, value string) {
	rdbValue := RDBValue{V: value, t: time.Now(), exp: -1}
	rdb[key] = &rdbValue
}

func SetEx(key string, value string, expt int64) {
	rdbValue := RDBValue{V: value, t: time.Now(), exp: expt}
	rdb[key] = &rdbValue
}

func Remove(key string) {
	delete(rdb, key)
}

func Get(key string) string {
	v, ok := rdb[key]
	if ok {
		if v.t.UnixMilli()+v.exp*1000 < time.Now().UnixMilli() {
			Remove(key)
			return ""
		}
		return v.V
	}
	return ""
}
