package gcs

import (

	// "strings"

	"fmt"
	"strings"
	"testing"

	"gitee.com/liumou_site/logger"
)

// 测试命令执行
func TestRunCmdBool(t *testing.T) {
	// 执行成功
	cc := ShellInit(true, true, true, "1")
	c := strings.Split("ls /home/liumou@ls /home/lll@lw /home/", "@")
	for _, s := range c {
		res, _, _ := cc.GetStatusOutputBool(s)
		if res {
			logger.Info("执行成功")
		} else {
			logger.Error("执行失败")
		}
	}
}

// 测试命令执行
func TestShell(t *testing.T) {
	// 执行成功
	cc := ShellInit(true, true, true, "1")
	c := strings.Split("ls /home/liumou@ls /home/lll@lw /home/", "@")
	for _, s := range c {
		res := cc.Shell(s)
		if res {
			logger.Info("执行成功")
		} else {
			logger.Error("执行失败")
		}
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

// 测试命令执行
func TestGetOut(t *testing.T) {
	// 执行成功
	cc := ShellInit(true, true, true, "1")
	c := strings.Split("ls /home/liumou@ls /home/lll@lw /home/@ls -la", "@")
	for _, s := range c {
		res, out, _ := cc.GetStatusOutputBool(s)
		if res {
			logger.Info("执行成功")
		} else {
			logger.Error("执行失败")
		}
		fmt.Println(out)
	}
}

// 测试实时刷新命令
func TestRe(t *testing.T) {
	// 执行成功
	c := strings.Split("ls /home/liumou@lss", "@")
	for _, s := range c {
		code := ShellSystem(s)
		if code != 0 {
			logger.Error("执行失败")
		} else {
			logger.Info("执行成功")
		}
		logger.Debug("退出代码: ", code)
	}
}
