package random

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"strings"
	"time"
)

func UUId() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func SmsCode(width int) string {
	nums := [10]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	l := len(nums)
	rand.NewSource(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%c", nums[rand.Intn(l)])
	}
	return sb.String()
}
