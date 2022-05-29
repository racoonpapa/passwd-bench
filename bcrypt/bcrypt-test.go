package main

import (
	"PasswdBench/util"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"time"
)

func getSaltedHash(password []byte, cost int) string {
	hash, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func bcryptTestByMultiCost(password string) {
	for cost := bcrypt.MinCost; cost < bcrypt.MaxCost; cost++ {
		before := time.Now()
		hash := getSaltedHash([]byte(password), cost)
		after := time.Now()

		if cost == bcrypt.DefaultCost {
			fmt.Printf("[%02d] Hash: %s, Duration: %v <-- Default\n", cost, hash, after.Sub(before))
		} else {
			fmt.Printf("[%02d] Hash: %s, Duration: %v\n", cost, hash, after.Sub(before))
		}
	}
}

func bcryptMatchTest(password string) time.Duration {
	fmt.Printf("Length: %02d - ", len(password))
	before := time.Now()
	hash := getSaltedHash([]byte(password), bcrypt.DefaultCost)
	after := time.Now()

	fmt.Printf("Hash: (%02d), Duration: %v - ", len(hash), after.Sub(before))

	before = time.Now()
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	after = time.Now()

	if err == nil {
		fmt.Printf("Matched. (%v)\n", after.Sub(before))
	} else {
		fmt.Printf("Not Matched. (%v)\n", after.Sub(before))
	}

	return after.Sub(before)
}

func main() {
	var password string
	rand.Seed(time.Now().Unix())

	//bcryptTestByMultiCost(getRandomString(32))
	for i := 8; i <= 40; i++ {
		password = util.GetRandomString(i)
		bcryptMatchTest(password)
	}
}
