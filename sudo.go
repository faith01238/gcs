package gcs

import (
	"fmt"

	"gitee.com/liumou_site/logger"
)

// 这是执行sudo命令的模块

// 通过Sudo权限执行命令并获取输出数据, 返回执行结果(成功/失败)及命令输出(字符串)
func (shells *ShellDebug) ShellGetOutSudo(cmd string) (bool, string) {
	if shells.is_root {
		logger.Debug("Currently running with root permission")
	} else {
		c := fmt.Sprintf("echo %s | sudo -S ", shells.password)
		cmd = c + cmd
		logger.Debug("Generate Command: ", cmd)
	}
	status, out, _ := shells.GetStatusOutputBool(cmd)
	return status, out
}

// 执行命令并获取执行结果, 返回执行结果(成功/失败)
func (shells *ShellDebug) ShellSudo(command string) bool {
	if shells.is_root {
		logger.Debug("Currently running with root permission")
	} else {
		c := fmt.Sprintf("echo %s | sudo -S ", shells.password)
		command = c + command
		logger.Debug("Generate Command: ", command)
	}
	return shells.Shell(command)
}
