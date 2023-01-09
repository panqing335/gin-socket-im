package util

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
)

func CheckPassword(password string, salt string, hash string) (bool bool, err error) {
	oldPwd := GenPwd(password, salt)
	bool = oldPwd == hash
	if !bool {
		err = errors.New("账户或密码错误，请重试")
	}
	return
}

func GenPwd(password string, salt string) string {
	pwd := []byte(password + salt)
	sum := sha1.Sum(pwd)
	toString := hex.EncodeToString(sum[:])

	return toString
}
