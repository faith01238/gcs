package gcs

// 定义命令的结构体
type ShellDebug struct {
	// // 是否显示Debug信息
	// debug bool
	// // 是否打印命令反馈信息
	// prints bool
	// // 是否打印错误输出
	printer, is_root, prints, debug bool
	password, user, home            string
	uid                             int
}
