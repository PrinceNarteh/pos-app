package utils

import (
	"crypto/rand"
	"fmt"
)

func SetCode(key string) string {
	code := rand.Text()
	return fmt.Sprintf("%s%s", key, code)
}

func SetOrderCode(key string) string {
	code := rand.Text()
	return fmt.Sprintf("%s%s", key, code)
}
