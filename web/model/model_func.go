package model

import (
	"my-micro/infra/imitate/vdb"
	"time"
)

var prefix = "img-verify-code:"

// 保存验证码到 数据库/redis
func SaveImgCode(uuid string, code string) error {
	vdb.SetEx(prefix+uuid, code, int64(time.Second*600))
	return nil
}

func CheckImgCode(uuid string, code string) bool {
	v := vdb.Get(prefix + uuid)
	if v == "" {
		return false
	}
	return v == code
}
