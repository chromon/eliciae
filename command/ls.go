package command

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 文件信息
type FileInf struct {
	mode os.FileMode
	count int64
	size int64
	date time.Time
	name string
	dir bool
}

type FileInfList []FileInf

// ls 命令
type Ls struct{
	all bool
	almostAll bool
	list bool
	time bool
	reverse bool
	recursive bool
	help bool
	helpStr string
}

// 初始化命令
func initLsCmd(ls *Ls) *flag.FlagSet {
	cmd := flag.NewFlagSet("ls", flag.ExitOnError)
	cmd.BoolVar(&ls.all, "a", false, "do not ignore entries starting with .")
	cmd.BoolVar(&ls.list, "l", false, "use a long listing format")
	cmd.BoolVar(&ls.time, "t", false, "sort by modification time, newest first")
	cmd.BoolVar(&ls.reverse, "r", false, "reverse order while sorting")
	cmd.BoolVar(&ls.help, "h", false, "display this help")

	// 帮助信息
	helpStr := `Usage: ls [OPTION]... [FILE]...
	List information about the FILEs (the current directory by default).
	Options:
  		-a do not ignore entries starting with .
		-l use a long listing format
		-t sort by modification time, newest first
		-r reverse order while sorting
		-h display this help and exit`

	ls.helpStr = helpStr
	return cmd
}

func LsCmd(cmdStr []string) error {

	ls := new (Ls)
	cmd := initLsCmd(ls)
	cmd.Parse(cmdStr[1:])

	// 查询文件列表全部信息
	//获取文件或目录相关信息
	fileList, errs := ioutil.ReadDir(CurrentPath)
	if errs != nil {
		return errs
	}

	fileInfList := queryFileInf(fileList)

	// 遍历实际传入的 flag
	var err error
	var returnVal bool
	var isList bool
	cmd.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "l":
			// ls -l 除了文件名之外，还将文件的权限、所有者、文件大小等信息详细列出来
			isList = true
		case "a":
			// ls -a 列出目录所有文件，包含以.开始的隐藏文件
			fileInfList = addAll(fileInfList)
		case "t":
			// ls -t 以文件修改时间排序
			sort.Slice(fileInfList, func(i, j int) bool {
				return fileInfList[i].date.Before(fileInfList[j].date)
			})
		case "r":
			// ls -r 反序排列
			sort.Slice(fileInfList, func(i, j int) bool {
				return fileInfList[i].name > fileInfList[j].name
			})
		case "h":
			_, err = fmt.Fprintln(os.Stdout, ls.helpStr)
			returnVal = true
		}
	})
	if err != nil {
		return err
	}

	traverse(fileInfList, isList)
	return nil
}

func queryFileInf(fileList []os.FileInfo) FileInfList {
	fileInfList := make(FileInfList, 0)
	for i := range fileList {
		fileInf := &FileInf{
			fileList[i].Mode(),
			1,
			fileList[i].Size(),
			fileList[i].ModTime(),
			fileList[i].Name(),
			fileList[i].IsDir(),
		}
		fileInfList = append(fileInfList, *fileInf)
	}
	return fileInfList
}

func addAll(fileInfList FileInfList) FileInfList {
	file1, _ := os.Stat(".")
	fileInf1 := &FileInf{
		file1.Mode(),
		1,
		file1.Size(),
		file1.ModTime(),
		file1.Name(),
		file1.IsDir(),
	}
	file2, _ := os.Stat("..")
	fileInf2 := &FileInf{
		file2.Mode(),
		1,
		file2.Size(),
		file2.ModTime(),
		file2.Name(),
		file2.IsDir(),
	}
	fileInfList = append([]FileInf{*fileInf2}, fileInfList...)
	fileInfList = append([]FileInf{*fileInf1}, fileInfList...)

	return fileInfList
}

func traverse(fileInfList FileInfList, isList bool) {

	for i := 0; i < len(fileInfList); i++ {
		var build strings.Builder
		if isList {
			build.WriteString(fileInfList[i].mode.String())
			build.WriteString("\t")
			build.WriteString(strconv.FormatInt(fileInfList[i].count, 10))
			build.WriteString("\t")
			build.WriteString(strconv.FormatInt(fileInfList[i].size, 10))
			build.WriteString("\t")
			build.WriteString(fileInfList[i].date.Format("Jan 02 15:04"))
			build.WriteString("\t")
		}

		build.WriteString(fileInfList[i].name)

		fmt.Fprintln(os.Stdout, build.String())
	}
}