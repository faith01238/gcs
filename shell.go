package gcs

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"gitee.com/liumou_site/logger"
	"github.com/bitfield/script"
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
	code, outs, err := shellbase(cmd, shells.debug)
	out := cast.ToString(outs)
	if code == 0 {
		return true, out, err
	} else {
		logger.Error(err)
		logger.Error("Execute command: ", cmd)
	}
	return false, out, err
}

// 使用第三方脚本二次开发
func (shells *ShellDebug) ScriptGetStatusOutput(cmds string) (excode int, out string, err error) {
	os_type := runtime.GOOS
	if strings.ToLower(os_type) == "windows" {
		cmd := exec.Command("cmd", "/C", cmds) //不加第一个第二个参数会报错
		if shells.prints {
			cmd.Stdout = os.Stdout // cmd.Stdout -> stdout  重定向到标准输出，逐行实时打印
		}
		if shells.printer {
			// 打印错误输出
			cmd.Stderr = os.Stderr // cmd.Stderr -> stderr
		}
		cmd.Stderr = os.Stderr // cmd.Stderr -> stderr
		//也可以重定向文件 cmd.Stderr= fd (文件打开的描述符即可)

		stdout, _ := cmd.StdoutPipe() //创建输出管道
		// defer stdout.Close()
		result, _ := io.ReadAll(stdout) // 读取输出结果
		resdata := string(result)
		cmd.Run()
		return cmd.ProcessState.ExitCode(), resdata, err
	} else {
		cmd := script.Exec(cmds)
		res, _ := cmd.String()
		code := cmd.ExitStatus()
		if shells.prints {
			logger.Debug("退出代码", code)
		}
		return code, res, cmd.Error()
	}

}

// 执行命令基础模块
func shellbase(cmd string, debug bool) (excode int, out string, err error) {
	os_type := runtime.GOOS
	if strings.ToLower(os_type) == "windows" {
		logger.Warn("暂不支持", os_type)
		return 0, "", fmt.Errorf("暂不支持当前系统")
	} else {
		// 创建命令实例
		run_cmd := exec.Command("bash", "-c", cmd)
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
