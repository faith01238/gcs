package gcs

import (
	"io/ioutil"
	"os"
	"os/exec"

	"gitee.com/liumou_site/logger"
	"github.com/spf13/cast"
)

// 执行命令并获取执行结果, 返回执行结果(成功/失败)
func (shells *ShellDebug) Shell(command string, name string) bool {
	logger.Info("Commands run: ", command)
	logger.Info("Task Name: ", name)
	code, _, err := shells.GetStatusOutput(command)
	if code == 0 {
		return true
	}
	logger.Error(err)
	return false
}

// 执行命令并获取输出数据, 返回执行结果布尔值 - (成功/失败)及 命令输出 (字符串)
func (shells *ShellDebug) GetStatusOutputBool(cmd, name string) (bool, string) {
	logger.Info(name)
	code, outs, err := shells.GetStatusOutput(cmd)
	out := cast.ToString(outs)
	if code == 0 {
		return true, out
	} else {
		logger.Error(err)
	}
	return false, err.Error()
}

// 获取命令状态码及输出
func (shells *ShellDebug) GetStatusOutput(cmds string) (excode int, out string, err error) {

	cmd := exec.Command("/bin/bash", "-c", cmds) //不加第一个第二个参数会报错
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
	defer stdout.Close()
	result, _ := ioutil.ReadAll(stdout) // 读取输出结果
	resdata := string(result)
	cmd.Run()
	return cmd.ProcessState.ExitCode(), resdata, err
}