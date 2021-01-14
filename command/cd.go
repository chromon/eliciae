package command

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func CdCmd(cmdStr []string) error {
	// 帮助信息
	helpStr := `Usage: cd [dir]
	Change the current directory to DIR.
	Options:
		-h  display this help and exit`

	if cmdStr[1] == "-h" {
		fmt.Fprintln(os.Stdout, helpStr)
		return nil
	}

	// 判断路径是否为目录
	isDir, err := checkDir(cmdStr[1])
	if err != nil {
		return err
	}

	// 非目录
	if !isDir {
		return errors.New("not a directory")
	}

	// 获取 path 的绝对路径
	CurrentPath, err = filepath.Abs(cmdStr[1])
	if err != nil {
		return err
	}
	// 切换当前 CurrentPath 为工作目录
	os.Chdir(CurrentPath)

	return err
}

// 判断是否是目录
func checkDir(path string) (bool, error) {
	file, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return file.IsDir(), nil
}