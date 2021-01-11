package command

import (
	"os"
)

// 当前目录路径
var CurrentPath string

// 命令分发
func DistributeCommand(cmdStr []string) error {
	switch cmdStr[0] {
	case "exit", "quit", "q":
		// $ exit / quit / q
		os.Exit(0)
	case "plus":
		// $ plus 2 3 4
		return Plus(cmdStr)
	case "cmd":
		// $ cmd dir /b
		return sysCommand("cmd", "/c", cmdStr[1:])
	case "bash":
		// $ bash ls -l
		return sysCommand("bash", "-c", cmdStr[1:])
	case "pwd":
		// $ pwd
		return Pwd(cmdStr)
	case "mkdir":
		// $ mkdir t
		// $ mkdir -m 0744 t
		// $ mkdir -p /tmp/t
		return MkdirCmd(cmdStr)
	case "rmdir":
		// $ rmdir a
		// $ rmdir -p ./a/b
		return RmdirCmd(cmdStr)
	}
	return nil
}