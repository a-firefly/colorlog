// root/colorlog/colorlog.go
package colorlog

import "fmt"

// 成功的日志样式
func Ok(format string, a ...any) {
	// \033 是八进制形式，十六进制是 0x1B
	format = "\033[0;0;32m" + format + "\033[0m\n"
	fmt.Printf(format, a...)
}
// 警告日志样式
func Warn(format string, a ...any) {
	format = "\033[0;0;33m" + format + "\033[0m\n"
	fmt.Printf(format, a...)
}
// 错误日志样式
func Error(format string, a ...any) {
	format = "\033[0;0;31m" + format + "\033[0m\n"
	fmt.Printf(format, a...)
}