package gcs

import (

	// "strings"

	"testing"

	"gitee.com/liumou_site/logger"
)

// 测试命令执行
func TestRunCmd(t *testing.T) {
	// 执行成功
	cc := ShellInit(true, true, true, "1")
	res, _ := cc.GetStatusOutputBool("dir")
	if res {
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
