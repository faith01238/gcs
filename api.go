package gcs

// 定义一个命令行密码结构体
type CommandPassword struct {
	password string
	user     string
	uid      int
	home     string
	is_root  bool
}

// 定义命令的结构体
type ShellDebug struct {
	// 是否显示Debug信息
	debug bool
	// 是否打印命令反馈信息
	prints bool
	// 是否打印错误输出
	printer bool
}
