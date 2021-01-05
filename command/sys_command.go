package command

import (
	"os"
	"os/exec"
	"strings"
)

// 调用 Windows / Linux 系统命令
func sysCommand(platform, sysArg string, args []string) error {
	cs := strings.Join(args, " ")
	cmd := exec.Command(platform, sysArg, cs)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
