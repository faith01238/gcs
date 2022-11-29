package gcs

// 定义命令的结构体
type ShellDebug struct {
	printer, is_root, prints, debug bool
	password, user, home            string
	uid                             int
}
