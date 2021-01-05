package command

import (
	"os"
)

// 命令分发
func DistributeCommand(cmdStr []string) error {
	switch cmdStr[0] {
	case "exit", "quit", "q":
		os.Exit(0)
	case "plus":
		return Plus(cmdStr)
	case "cmd":
		return sysCommand("cmd", "/c", cmdStr[1:])
	case "bash":
		return sysCommand("bash", "-c", cmdStr[1:])
	}
	return nil
}