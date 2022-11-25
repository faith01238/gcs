package gcs

import (
	"fmt"

	// "strings"
	"testing"

	"gitee.com/liumou_site/logger"
)

// // 测试命令执行
func TestRunSudo(t *testing.T) {
	sudo := SudoCommand("11")
	sudo_c := sudo.ShellSudo("ls /home", "Sudo测试")
	if sudo_c {
		fmt.Println("sudo命令执行成功")
	} else {
		fmt.Println("Sudo命令执行失败")
	}
}

// 测试命令执行
func TestRunCmd(t *testing.T) {
	cc := ShellInit(true, true, true)
	res := cc.Shell("ls", "ls")
	if res {
		logger.Info("执行成功")
	} else {
		logger.Error("执行失败")
	}
}
