package main

import (
	"PasswdBench/util"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"math/rand"
	"time"
)

func benchmark(password string, iter int) {
	before := time.Now()
	salt := util.GetRandomString(32)
	_ = pbkdf2.Key([]byte(password), []byte(salt), iter, 60, sha1.New)
	after := time.Now()
	fmt.Printf("[%d] %d Length password hash: %v\n", iter, len(password), after.Sub(before))
}

func benchmarkByLength(length int) {
	password := util.GetRandomString(length)
	iter := 100000

	before := time.Now()
	salt := util.GetRandomString(32)
	_ = pbkdf2.Key([]byte(password), []byte(salt), iter, 60, sha256.New)
	after := time.Now()
	fmt.Printf("[%d] %d Length password hash: %v\n", iter, len(password), after.Sub(before))
}

func main() {
	var iters = []int{512, 1024, 4096, 10000, 100000, 1000000}
	var password string
	rand.Seed(time.Now().Unix())

	password = util.GetRandomString(32)

	for i := 0; i < len(iters); i++ {
		benchmark(password, iters[i])
	}

	for i := 8; i <= 40; i++ {
		benchmarkByLength(i)
	}
}
