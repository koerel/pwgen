package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(len int, charset string) string {
	start := 97
	if charset == "full" {
		start = 33
	}
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(start, 122))
	}
	return string(bytes)
}

func main() {
	args := os.Args[1:]
	rand.Seed(time.Now().UnixNano())
	l, _ := strconv.Atoi(args[0])
	charset := "small"
	if len(args) > 1 {
		charset = args[1]
	}
	fmt.Println(randomString(l, charset))
}
