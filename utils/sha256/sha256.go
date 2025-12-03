package sha256

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"hash"
)

// SHA256Crypto SHA256 加密器
type SHA256Crypto struct {
	hash hash.Hash
}

// NewSHA256Crypto 创建新的SHA256加密器
func NewSHA256Crypto() *SHA256Crypto {
	return &SHA256Crypto{
		hash: sha256.New(),
	}
}

// Hash 计算字符串的SHA256哈希值
func (h *SHA256Crypto) Hash(data string) string {
	h.hash.Write([]byte(data))
	return hex.EncodeToString(h.hash.Sum(nil))
}

// HashBytes 计算字节切片的SHA256哈希值
func (h *SHA256Crypto) HashBytes(data []byte) string {
	h.hash.Write(data)
	return hex.EncodeToString(h.hash.Sum(nil))
}

// HashToBase64 计算SHA256并返回Base64编码
func (h *SHA256Crypto) HashToBase64(data string) string {
	h.hash.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.hash.Sum(nil))
}
