package utils

import (
	"math/rand"
	"time"

	"gopkg.in/night-codes/types.v1"
)

func GenerateToken() string {
	rand.Seed(time.Now().UnixNano())
	return types.String(rand.Int())
}
