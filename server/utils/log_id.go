package utils

import (
	"fmt"
	"time"
)

func NewLogId() string {
	now := time.Now()
	fmt.Println(0x100000)
	return fmt.Sprintf("%08x%08x", now.Unix(), now.UnixNano()%0x1e6)
}
