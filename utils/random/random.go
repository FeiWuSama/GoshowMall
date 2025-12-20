package random

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"strings"
	"time"
)

func GenUUId() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func GenUserUUId() int64 {
	return int64(uuid.New().ID())
}

func GenSmsCode(width int) string {
	nums := [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	l := len(nums)
	rand.NewSource(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%c", nums[rand.Intn(l)])
	}
	return sb.String()
}
