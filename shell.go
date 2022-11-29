package gcs

import (
	"fmt"
	"os/exec"

	"gitee.com/liumou_site/logger"
	"github.com/spf13/cast"
)

// 执行命令并获取执行结果, 返回执行结果(成功/失败)
func (shells *ShellDebug) Shell(command string) bool {
	if shells.prints {
		logger.Info("Commands run: ", command)
	}
	code, out := system(command, false)
	if shells.prints {
		fmt.Println(out)
	}
	if code != 0 {
		return false
	}
	return true
}

// 执行命令并获取输出数据, 返回执行结果布尔值 - (成功/失败)及 命令输出 (字符串)
func (shells *ShellDebug) GetStatusOutputBool(cmd string) (bool, string) {
	code, out := system(cmd, shells.debug)
	if code == 0 {
		return true, out
	}
	return false, out
}

// 执行命令基础模块
// func shellbase(cmd string, debug bool) (excode int, out string, err error) {
// 	os_type := runtime.GOOS
// 	if strings.ToLower(os_type) == "windows" {
// 		logger.Warn("暂不支持", os_type)
// 		return 0, "", fmt.Errorf("暂不支持当前系统")
// 	} else {
// 		// 创建命令实例
// 		run_cmd := exec.Command("/bin/bash", "-c", cmd)
// 		// 开始运行(使用堵塞方式)
// 		err := run_cmd.Run()
// 		// 获取执行状态码
// 		exit_code := run_cmd.ProcessState.ExitCode()
// 		// 获取输出
// 		buf, _ := run_cmd.CombinedOutput() // 错误处理略
// 		get_out := string(buf)
// 		if debug {
// 			logger.Debug("Exit Code: ", exit_code)
// 			logger.Debug("Output information: ", get_out)
// 		}
// 		// 将字节转换到字符串
// 		res := cast.ToString(get_out)
// 		// 返回信息
// 		return exit_code, res, err
// 	}

// }

// 执行命令并实时获取输出,和Python的os.system函数一样,返回退出代码
// (excode int, out string, err error)
func ShellSystem(command string) (code int) {
	code, _ = system(command, false)
	return code
}

// 执行命令并支持实时获取输出,和Python的os.system函数一样,返回退出代码
// command 需要执行的命令  retime 是否开启实时打印
func system(command string, retime bool) (code int, out string) {
	cmd := exec.Command("/bin/bash", "-c", command)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		logger.Error("管道创建失败")
		logger.Debug("Execute command: ", command)
		return 1, "管道创建失败"
	}
	run_err := cmd.Start()
	if run_err != nil {
		logger.Error("命令启动失败 :", err)
		logger.Debug("Execute command: ", command)
		return 2, "命令启动失败"
	}
	// 创建字符串保存输出信息
	var result string
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		// 如果开启了实时打印,则将信息逐行输出到终端
		if retime {
			fmt.Print(string(tmp))
		}
		result = result + cast.ToString(string(tmp))
		if err != nil {
			break
		}
	}
	// 等待命令退出，并等待任何复制到stdin或从stdout或stderr复制完成。
	wait_err := cmd.Wait()
	// ExitCode返回已退出进程的退出代码，如果进程尚未退出或被信号终止，则返回-1。
	exit_code := cmd.ProcessState.ExitCode()
	if exit_code != 0 {
		logger.Error("Error :", wait_err)
		logger.Debug("Execute command: ", command)
		return exit_code, result
	}
	return exit_code, result
}
