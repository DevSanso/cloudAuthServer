package utils

import (
	"crypto/sha256"

)

var salt = RandStringBytes(8)


func Make256Hash(val ...string) []byte {
	sha := sha256.New()
	for _,v := range val {
		sha.Write([]byte(v))
	}

	return sha.Sum([]byte(""))
}

func Make256HashSalt(val ...string) []byte {
	v := append(make([]string,0),salt)
	v = append(v,val...)
	return Make256Hash(v...)
}