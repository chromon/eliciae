package command

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// 加法求和
func Plus(cmdStr []string) error {
	if len(cmdStr) < 3 {
		return errors.New("required for 2 arguments")
	}

	var arrNum []int64
	for _, arg := range cmdStr[1:] {
		n, _ := strconv.ParseInt(arg, 10, 64)
		arrNum = append(arrNum, n)
	}
	fmt.Fprintln(os.Stdout, sum(arrNum...))

	return nil
}

func sum(numbers ...int64) int64 {
	res := int64(0)
	for _, num := range numbers {
		res += num
	}
	return res
}
