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
