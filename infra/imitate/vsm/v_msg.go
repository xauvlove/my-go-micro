package vsm

import (
	"fmt"
	"math/rand"
)

func GenVerifyCode(phone string) (string, error) {
	// 生成 6 位随机数
	code := fmt.Sprintf("%06d", rand.Int31n(1000000))
	return code, nil
}
