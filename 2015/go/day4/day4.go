package day4

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func Run() {
	secret := "ckczppom"
	p1 := part1(secret)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(secret)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(secret string) int {
	for i := 1; ; i++ {
		s := fmt.Sprintf("%v%v", secret, i)
		hash := md5.Sum([]byte(s))
		if strings.HasPrefix(hex.EncodeToString(hash[:]), "00000") {
			return i
		}
	}
}

func part2(secret string) int {
	for i := 1; ; i++ {
		s := fmt.Sprintf("%v%v", secret, i)
		hash := md5.Sum([]byte(s))
		if strings.HasPrefix(hex.EncodeToString(hash[:]), "000000") {
			return i
		}
	}
}
