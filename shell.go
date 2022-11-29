package gcs

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"gitee.com/liumou_site/logger"
	"github.com/spf13/cast"
)

// 执行命令并获取执行结果, 返回执行结果(成功/失败)
func (shells *ShellDebug) Shell(command string) bool {
	if shells.prints {
		logger.Info("Commands run: ", command)
	}
	status, out, err := shells.GetStatusOutputBool(command)
	if !status {
		logger.Error(err)
	}
	if shells.prints {
		fmt.Println(out)
	}
	return status
}

// 执行命令并获取输出数据, 返回执行结果布尔值 - (成功/失败)及 命令输出 (字符串)
func (shells *ShellDebug) GetStatusOutputBool(cmd string) (bool, string, error) {
	// code, outs, err := shells.ScriptGetStatusOutput(cmd)
	code, out, err := shellbase(cmd, shells.debug)
	if code == 0 {
		return true, out, err
	} else {
		logger.Error(err)
		logger.Error("Execute command: ", cmd)
	}
	return false, out, err
}

// 执行命令基础模块
func shellbase(cmd string, debug bool) (excode int, out string, err error) {
	os_type := runtime.GOOS
	if strings.ToLower(os_type) == "windows" {
		logger.Warn("暂不支持", os_type)
		return 0, "", fmt.Errorf("暂不支持当前系统")
	} else {
		// 创建命令实例
		run_cmd := exec.Command("/bin/bash", "-c", cmd)
		// 开始运行(使用堵塞方式)
		err := run_cmd.Run()
		// 获取执行状态码
		exit_code := run_cmd.ProcessState.ExitCode()
		// 获取输出
		buf, _ := run_cmd.CombinedOutput() // 错误处理略
		get_out := string(buf)
		if debug {
			logger.Debug("Exit Code: ", exit_code)
			logger.Debug("Output information: ", get_out)
		}
		// 将字节转换到字符串
		res := cast.ToString(get_out)
		// 返回信息
		return exit_code, res, err
	}

}

// 执行命令并实时获取输出
// (excode int, out string, err error)
func ShellSystem(command string) (code int) {
	cmd := exec.Command("/bin/bash", "-c", command)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		exit_code := cmd.ProcessState.ExitCode()
		logger.Error("Error :", err)
		logger.Debug("Execute command: ", command)
		return exit_code
	}
	run_err := cmd.Start()
	if run_err != nil {
		exit_code := cmd.ProcessState.ExitCode()
		logger.Error("Error :", err)
		logger.Debug("Execute command: ", command)
		return exit_code
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}
	wait_err := cmd.Wait()
	exit_code := cmd.ProcessState.ExitCode()
	if exit_code != 0 {
		logger.Error(wait_err)
		logger.Error("Error :", err)
		logger.Debug("Execute command: ", command)
		return exit_code
	}
	return exit_code
}
