package model

import (
	"crypto/md5"
	"encoding/hex"
	"my-micro/infra/imitate/vdb"
	"time"
)

var imgPrefix = "img-verify-code:"
var smsPrefix = "sms-verify-code:"

// 保存验证码到 数据库/redis
func SaveImgCode(uuid string, code string) error {
	vdb.SetEx(imgPrefix+uuid, code, int64(time.Second*600))
	return nil
}

func CheckImgCode(uuid string, code string) bool {
	v := vdb.Get(imgPrefix + uuid)
	if v == "" {
		return false
	}
	return v == code
}

func SaveSmsCode(phone string, code string) error {
	vdb.SetEx(smsPrefix+phone, code, int64(time.Minute*10))
	return nil
}

func CheckSmsCode(phone string, code string) bool {
	v := vdb.Get(smsPrefix + phone)
	if v == "" {
		return false
	}
	return v == code
}

func RegisterUser(phone string, password string) error {
	var user User
	// 使用手机号作为用户名
	user.Name = phone

	m5 := md5.New()
	m5.Write([]byte(password))
	pwdHash := hex.EncodeToString(m5.Sum(nil))
	user.Password_hash = pwdHash

	// insert
	return GlobalConn.Create(&user).Error
}
