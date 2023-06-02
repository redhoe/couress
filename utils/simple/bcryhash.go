package simple

import "github.com/jameskeane/bcrypt"

func DecodeHashBySignKey(str, slat string) (string, error) {
	return bcrypt.Hash(str, slat)
}

func CheckSignHash(str, slat, hashStr string) bool {
	nowStr, err := bcrypt.Hash(str, slat)
	if err != nil {
		return false
	}
	return nowStr == hashStr
}
