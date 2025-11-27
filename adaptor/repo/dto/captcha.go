package dto

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type CaptchaDto struct {
	Once string `url:"once"`
	Time int64  `url:"ts"`
	Sign string `url:"sign"`
}

func (c *CaptchaDto) CheckSign() bool {
	data := fmt.Sprintf("%s%s%d", c.Once, "fewiwusama2015", c.Time)
	hash := sha256.Sum256([]byte(data))
	return c.Sign == hex.EncodeToString(hash[:])
}
