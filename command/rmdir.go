package command

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// rmdir 相关信息
type Rmdir struct {
	parents bool
	directory string
	help bool
	helpStr string
}

// 初始化命令
func initRmdirCmd(rmdir *Rmdir) *flag.FlagSet {
	cmd := flag.NewFlagSet("mkdir", flag.ExitOnError)
	cmd.BoolVar(&rmdir.parents, "p", false, "-p, --parents remove DIRECTORY and its ancestors; e.g., 'rmdir -p a/b/c'")
	cmd.BoolVar(&rmdir.help, "h", false, "display this help")

	// 帮助信息
	helpStr := `Usage: rmdir [OPTION]... DIRECTORY...
	Remove the DIRECTORY(ies), if they are empty.
	Options:
  		-p, --parents   remove DIRECTORY and its ancestors; e.g., 'rmdir -p a/b/c'
		-h  display this help and exit`
	rmdir.helpStr = helpStr

	return cmd
}

func RmdirCmd(cmdStr []string) error {
	// 构建命令并初始化
	rmdir := new (Rmdir)
	cmd := initRmdirCmd(rmdir)
	cmd.Parse(cmdStr[1:])

	// 遍历实际传入的 flag
	var err error
	var returnVal bool
	cmd.Visit(func(f *flag.Flag) {
		if f.Name == "h" {
			fmt.Fprintln(os.Stdout, rmdir.helpStr)
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

	// 目录参数
	if len(cmd.Args()) > 0 {
		rmdir.directory = cmd.Args()[0]
	} else {
		return errors.New("rmdir: missing operand\nTry 'rmdir -h' for more information.")
	}

	if rmdir.parents {
		// 删除多级目录
		return removeDir(CurrentPath + string(os.PathSeparator) + rmdir.directory)
	} else {
		// 删除单一文件夹
		// 判断目录是否为空
		info, err := ioutil.ReadDir(CurrentPath + string(os.PathSeparator) + rmdir.directory)
		if err != nil {
			// 目录不存在
			return err
		}

		if len(info) == 0 {
			// 目录为空，删除
			os.RemoveAll(CurrentPath + string(os.PathSeparator) + rmdir.directory)
		} else {
			// 目录不为空无法删除
			errStr := fmt.Sprintf("rmdir: failed to remove '%s': Directory not empty", rmdir.directory)
			return errors.New(errStr)
		}
	}

	return nil
}

// 递归删除空目录
func removeDir(path string) error {
	// 判断目录是否为空
	info, err := ioutil.ReadDir(path)
	if err != nil {
		// 目录不存在
		return err
	}

	if len(info) == 0 {
		// 目录为空，删除
		os.RemoveAll(path)

		// 判断是否是最终目录
		s := filepath.Dir(path)
		if s == CurrentPath {
			return nil
		}

		return removeDir(filepath.Dir(path))
	} else {
		// 目录不为空无法删除
		errStr := fmt.Sprintf("rmdir: failed to remove '%s': Directory not empty", path)
		return errors.New(errStr)
	}
}
