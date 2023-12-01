package vsm

import (
	"fmt"
	"math/rand"
	"my-micro/infra/imitate/vdb"
	"time"
)

var prefix = "sms-verify-code:"

func GenVerifyCode(phone string) (string, error) {
	// 生成 6 位随机数
	code := fmt.Sprintf("%06d", rand.Int31n(1000000))
	vdb.SetEx(prefix+phone, code, int64(time.Minute*10))
	return code, nil
}
