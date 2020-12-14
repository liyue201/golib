package xhash

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func Md5(data []byte) string {
	h := md5.New()
	if _, err := h.Write(data); err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

func Sha256(data []byte) string {
	h := sha256.New()
	if _, err := h.Write(data); err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}
