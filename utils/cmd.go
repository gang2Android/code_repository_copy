package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

func GitClone(basePath string, repository string) {
	fmt.Println(basePath, repository)
	repositoryName := GetFileName(repository)

	b, _ := PathExists(basePath)
	if b == false {
		err := os.MkdirAll(basePath, os.ModePerm) //在当前目录下生成md目录
		if err != nil {
			fmt.Println(err)
		}
	}

	var cmdStr strings.Builder
	cmdStr.WriteString("cd " + basePath)
	cmdStr.WriteString(" && ")
	cmdStr.WriteString("git clone " + repository)
	cmdStr.WriteString(" && ")
	cmdStr.WriteString("cd " + repositoryName)

	fmt.Println(cmdStr.String())

	Cmd(cmdStr.String())
}

func Cmd(cmdStr string) {
	fmt.Println("Cmd=", cmdStr)
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		cmd = exec.Command("sh", "-c", cmdStr)
	}
	stdout, _ := cmd.StdoutPipe()
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		fmt.Printf("cmd.Start: %v")
	}
	result, _ := ioutil.ReadAll(stdout)
	resdata := string(result)
	var res int
	if err := cmd.Wait(); err != nil {
		if ex, ok := err.(*exec.ExitError); ok {
			fmt.Println("cmd exit status")
			res = ex.Sys().(syscall.WaitStatus).ExitStatus() //获取命令执行返回状态，相当于shell: echo $?
		}
	}
	fmt.Println("Cmd-result=", resdata, res)

}

func GetFileName(filePath string) string {
	n := strings.LastIndex(filePath, "/") + 1
	if n == -1 {
		n = 0
	}
	filePath = string([]byte(filePath)[n:])
	m := strings.Index(filePath, ".")
	if m == -1 {
		m = len(filePath)
	}
	filePath = string([]byte(filePath)[:m])
	return filePath
}
