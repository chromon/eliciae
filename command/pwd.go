package command

import (
	"flag"
	"fmt"
	"os"
)

// 查看当前工作目录路径
func Pwd(cmdStr []string) error {
	pwd := flag.NewFlagSet("pwd", flag.ExitOnError)
	pwd.Bool("h", false, "display this help")
	err := pwd.Parse(cmdStr[1:])

	// 帮助信息
	helpStr := `Usage: pwd [-h]
	Print the name of the current working directory.
	Options:
		-h	display this help and exit`

	var returnVal bool
	pwd.Visit(func(f *flag.Flag) {
		if f.Name == "h" {
			_, err = fmt.Fprintln(os.Stdout, helpStr)
			returnVal = true
		}
	})
	if err != nil {
		return err
	}
	// 仅输出帮助信息后返回
	if returnVal {
		return nil
	}

	// 获取当前目录
	//CurrentPath, err := os.Getwd()
	_, err = fmt.Fprintln(os.Stdout, CurrentPath)
	return err
}