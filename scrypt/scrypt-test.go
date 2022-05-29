package main

import (
	"PasswdBench/util"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"log"
	"math"
	"time"
)

func benchmark(pow int) {
	password := util.GetRandomString(32)

	N := int(math.Pow(2, float64(pow)))

	before := time.Now()
	salt := util.GetRandomString(32)
	dk, err := scrypt.Key([]byte(password), []byte(salt), N, 8, 1, 32)
	if err != nil {
		log.Println(err)
		return
	}
	result := base64.StdEncoding.EncodeToString(dk)
	after := time.Now()

	fmt.Printf("%v (length: %v/N = %v/ %v)\n", result, len(result), N, after.Sub(before))
}

func main() {
	for i := 10; i <= 20; i++ {
		benchmark(i)
	}
}
