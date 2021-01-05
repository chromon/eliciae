package command

import (
	"fmt"
	"os"
)

// 查看当前工作目录路径
func Pwd() error {
	CurrentPath, err := os.Getwd()
	fmt.Fprintln(os.Stdout, CurrentPath)
	return err
}
