package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// 基础AES加密函数（CBC模式）
func EncryptAES(key, plaintext []byte) ([]byte, error) {
	// 1. 创建AES密码块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 2. 填充数据（PKCS7填充）
	plaintext = pkcs7Padding(plaintext, aes.BlockSize)

	// 3. 生成随机IV
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// 4. 加密
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

// 基础AES解密函数（CBC模式）
func DecryptAES(key, ciphertext []byte) ([]byte, error) {
	// 1. 创建AES密码块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 2. 检查密文长度
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("密文太短")
	}

	// 3. 分离IV和密文
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// 4. 解密
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// 5. 移除填充
	return pkcs7UnPadding(ciphertext)
}

// PKCS7填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// PKCS7去填充
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("数据为空")
	}
	unPadding := int(data[length-1])
	if unPadding > length {
		return nil, errors.New("填充错误")
	}
	return data[:(length - unPadding)], nil
}
