package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

func genRandString(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func main() {
	// generate and store random string on startup
	randomStr, err := genRandString(12)
	if err != nil {
		fmt.Printf("%s error generating random string: %v\n", time.Now().Format(time.RFC3339), err)
		return
	}

	// print once immediately
	fmt.Printf("%s %s\n", time.Now().Format(time.RFC3339), randomStr)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for t := range ticker.C {
		fmt.Printf("%s %s\n", t.Format(time.RFC3339), randomStr)
	}
}
