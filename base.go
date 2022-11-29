package gcs

import (
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"strconv"
)

// 这是创建sudo命令实例的构造函数

func SudoCommand(password string) *CommandPassword {
	c := new(CommandPassword)
	c.password = password
	get, username, uid, u_home := GetUserInfo(false)
	if get {
		c.home = u_home
		c.user = username
		c.uid = uid
	} else {
		c.home = "None"
		c.user = "None"
		c.uid = 10000
	}
	if c.user == "root" {
		c.is_root = true
	}
	return c
}

// 命令实例构造函数
func ShellInit(debug, prints, ptinter bool) *ShellDebug {
	shell := new(ShellDebug)
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
