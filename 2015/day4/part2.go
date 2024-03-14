package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	input := "ckczppom"
	nonce := 0

	for {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", input, nonce)))
		hashString := fmt.Sprintf("%s\r", hex.EncodeToString(hash[:]))

		if strings.HasPrefix(hashString, "000000") {
			break
		}
		nonce += 1
	}
	fmt.Println(nonce)
}
