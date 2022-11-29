package gcs

import (

	// "strings"
	"exec"
	"os/exec"
	"testing"

	"gitee.com/liumou_site/logger"
)

// // 测试命令执行
// func TestRunSudo(t *testing.T) {
// 	sudo := SudoCommand("11")
// 	sudo_c := sudo.ShellSudo("ls /home", "Sudo测试")
// 	if sudo_c {
// 		fmt.Println("sudo命令执行成功")
// 	} else {
// 		fmt.Println("Sudo命令执行失败")
// 	}
// }

// 测试命令执行
func TestRunCmd(t *testing.T) {
	// 执行成功
	cc := ShellInit(true, true, true)
	res, _, _ := cc.Shell("ls /home", "列出用户目录")
	if res == 0 {
		logger.Info("执行成功")
	} else {
		logger.Error("执行失败")
	}
	i := exec.LookPath("ls")
	// 执行失败
	ress, _, _ := cc.Shell("ipconfisg")
	if ress == 0 {
		logger.Info("执行成功")
	} else {
		logger.Error("执行失败")
	}
}

func TestCheckCmd(t *testing.T) {
	ls_ := CheckCmd("dir")
	if ls_ {
		logger.Info("存在")
	} else {
		logger.Error("不存在命令")
	}
}
