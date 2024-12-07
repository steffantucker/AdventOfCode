package utils

import (
	"fmt"
	"strconv"
)

func MustAtoi(n string) int {
	i, err := strconv.Atoi(n)
	if err != nil {
		panic(fmt.Sprintf("Must failed, provided: %#v, got error: %v", n, err))
	}
	return i
}

func MustAtoi64(n string) int64 {
	i, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Must failed, provided: %#v was not parsed. %v", n, err))
	}
	return i
}
