package main

import (
	"bufio"
	"eliciae/command"
	"fmt"
	"os"
	"strings"
)

func main() {

	// 当前目录路径
	command.CurrentPath, _ = os.Getwd()

	// 读取输入
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(command.CurrentPath, ":$ ")
		cmdStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// 运行命令
		err = runCommand(cmdStr)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// 执行命令
func runCommand(cmdStr string) error {

	cmdStr = strings.TrimSuffix(cmdStr, "\r\n");
	commandArray := strings.Fields(cmdStr)

	return command.DistributeCommand(commandArray)
}

