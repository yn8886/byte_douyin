package utils

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	pwd := "123456"
	pwd, _ = PwdHash(pwd)
	fmt.Println(pwd)
}