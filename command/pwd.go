package command

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
)

// 查看当前工作目录路径
func Pwd(cmdStr []string) error {

	// 帮助信息
	helpStr := `Usage: pwd [-h]
	Print the name of the current working directory.
	Options:
		-h	display this help and exit
	`

	if len(cmdStr) == 1 {
		//CurrentPath, err := os.Getwd()
		_, err := fmt.Fprintln(os.Stdout, CurrentPath)
		return err
	} else if len(cmdStr) == 2 && cmdStr[1] == "-h" {
		_, err := fmt.Fprintln(os.Stdout, helpStr)
		return err
	}

	// error 信息
	errStr := "bad command syntax: "

	var buffer bytes.Buffer
	buffer.WriteString(errStr)
	buffer.WriteString(strings.Join(cmdStr, " "))
	buffer.WriteString("\n")
	buffer.WriteString(helpStr)

	return errors.New(buffer.String())
}
