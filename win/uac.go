package win

import (
	"golang.org/x/sys/windows"
	"os"
	"strings"
	"syscall"
)

//Request UAC elevation in Go,要求以管理员身份运行程序
func RequireUAC() error {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		verb := "runas"
		exe, _ := os.Executable()
		cwd, _ := os.Getwd()
		args := strings.Join(os.Args[1:], " ")

		verbPtr, _ := syscall.UTF16PtrFromString(verb)
		exePtr, _ := syscall.UTF16PtrFromString(exe)
		cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
		argPtr, _ := syscall.UTF16PtrFromString(args)

		var showCmd int32 = 1 //SW_NORMAL

		err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
		return err
	}
	return nil
}
