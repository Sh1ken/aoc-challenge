package year2015

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

type Day04 struct{}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func searchAdjacentNumber(secretKey string, prefix string) int {
	number := 1

	for {
		hash := getMD5Hash(secretKey + strconv.Itoa(number))

		if strings.HasPrefix(hash, prefix) {
			return number
		}

		number++
	}
}

func (p Day04) PartA(lines []string) any {
	secretKey := lines[0]
	return searchAdjacentNumber(secretKey, "00000")
}

func (p Day04) PartB(lines []string) any {
	secretKey := lines[0]
	return searchAdjacentNumber(secretKey, "000000")
}
