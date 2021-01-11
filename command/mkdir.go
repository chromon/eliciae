package command

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

// mkdir 相关命令信息
type Mkdir struct {
	mode uint64
	parents bool
	directory string
	help bool
	helpStr string
}

// 初始化命令
func initCmd(mkdir *Mkdir) *flag.FlagSet {
	cmd := flag.NewFlagSet("mkdir", flag.ExitOnError)
	cmd.Uint64Var(&mkdir.mode, "m", 0777, "-m, --mode=MODE set file mode (as in chmod), not a=rwx - umask")
	cmd.BoolVar(&mkdir.parents, "p", false, "-p, --parents no error if existing, make parent directories as needed")
	cmd.BoolVar(&mkdir.help, "h", false, "display this help")

	// 帮助信息
	helpStr := `Usage: mkdir [OPTION]... DIRECTORY...
	Create the DIRECTORY(ies), if they do not already exist.
	Options:
  		-m  set file mode (as in chmod), not a=rwx - umask
  		-p  no error if existing, make parent directories as needed
		-h  display this help and exit`
	mkdir.helpStr = helpStr

	return cmd
}

func MkdirCmd(cmdStr []string) error {

	// 构建命令并初始化
	mkdir := new (Mkdir)
	cmd := initCmd(mkdir)
	cmd.Parse(cmdStr[1:])

	// 遍历实际传入的 flag
	var err error
	var returnVal bool
	cmd.Visit(func(f *flag.Flag) {
		if f.Name == "m" {
			mkdir.mode, err = strconv.ParseUint(f.Value.String(), 10, 32)
		} else if f.Name == "h" {
			fmt.Fprintln(os.Stdout, mkdir.helpStr)
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
		mkdir.directory = cmd.Args()[0]
	} else {
		return errors.New("mkdir: missing operand\nTry 'mkdir -h' for more information.")
	}

	if mkdir.parents {
		// 创建多级目录
		err = os.MkdirAll(CurrentPath + string(os.PathSeparator) + mkdir.directory, os.FileMode(mkdir.mode))
		if err != nil {
			return err
		}
	} else {
		// 创建单一文件夹
		err = os.Mkdir(CurrentPath + string(os.PathSeparator) + mkdir.directory, os.FileMode(mkdir.mode))
		if err != nil {
			return err
		}
	}

	return nil
}
