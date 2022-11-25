package gcs

import (
	"fmt"

	"gitee.com/liumou_site/logger"
)

// 这是执行sudo命令的模块

// 通过Sudo权限执行命令并获取输出数据, 返回执行结果(成功/失败)及命令输出(字符串)
func (p *CommandPassword) ShellGetOutSudo(cmd, name string) (bool, string) {
	run := ShellInit(true, true, true)
	logger.Info(name)
	if p.is_root {
		logger.Debug("Currently running with root permission")
	} else {
		c := fmt.Sprintf("echo %s | sudo -S ", p.password)
		cmd = c + cmd
		logger.Debug("Generate Command: ", cmd)
	}
	return run.GetStatusOutputBool(cmd, name)
}

// 执行命令并获取执行结果, 返回执行结果(成功/失败)
func (p *CommandPassword) ShellSudo(command, name string) bool {
	run := ShellInit(true, true, true)
	if p.is_root {
		logger.Debug("Currently running with root permission")
	} else {
		c := fmt.Sprintf("echo %s | sudo -S ", p.password)
		command = c + command
		logger.Debug("Generate Command: ", command)
	}
	return run.Shell(command, name)
}
