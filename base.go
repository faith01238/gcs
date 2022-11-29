package gcs

import (
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"strconv"
)

// 命令实例构造函数, 当不需要执行sudo命令的时候直接传入任意字符串给password即可
func ShellInit(debug, prints, ptinter bool, password string) *ShellDebug {
	shell := new(ShellDebug)
	// 系统信息
	shell.password = password
	get, username, uid, u_home := GetUserInfo(false)
	if get {
		shell.home = u_home
		shell.user = username
		shell.uid = uid
	} else {
		shell.home = "None"
		shell.user = "None"
		shell.uid = 10000
	}
	if shell.user == "root" {
		shell.is_root = true
	}
	// 调试开关
	shell.debug = debug
	shell.printer = ptinter
	shell.prints = prints
	return shell
}

func GetUserInfo(display bool) (ok bool, username string, userid int, UserHome string) {
	var id int
	// fmt.Println(path)
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
		return false, "None", 10000, "None"
	}
	name := currentUser.Name
	id, err = strconv.Atoi(currentUser.Uid)
	if err != nil {
		id = 10000
	}
	home := currentUser.HomeDir
	if display {
		fmt.Println("UserName is: ", name)
		fmt.Println("UserId is: ", id)
		fmt.Println("UserHome : ", home)
	}
	return true, name, id, home
}

// 检查命令是否存在
func CheckCmd(cmd string) bool {
	// 通过此函数可以从path变量查询命令是否存在，返回命令绝对路径和查找结果
	_, err := exec.LookPath(cmd)
	return err == nil
}
