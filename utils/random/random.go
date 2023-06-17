package random

import (
	"math/rand"
	"time"
)

func RandInt(min, max int) int {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    return r.Intn(max-min) + min
}

func RandFloat(min, max float64) float64 {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    return min + r.Float64()*(max-min)
}