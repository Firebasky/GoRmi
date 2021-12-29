package utils

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func PrintLogo(version, author string) {
	logo := "                                   \n   mmm         mmmmm           \"   \n m\"   \"  mmm   #   \"# mmmmm  mmm   \n #   mm #\" \"#  #mmmm\" # # #    #   \n #    # #   #  #   \"m # # #    #   \n  \"mmm\" \"#m#\"  #    \" # # #  mm#mm \n"
	fmt.Println(logo)
	fmt.Printf("version: %s   author: %s\n", version, author)
}
func Info(format string, a ...interface{}) {
	now := getTime()
	var data string
	if checkOS() {
		data = fmt.Sprintf("[+] [%s] %s\n", now, format)
	} else {
		data = fmt.Sprintf("\x1b[32m[+] [%s] %s\x1b[0m\n", now, format)
	}
	_, _ = fmt.Fprintf(os.Stdout, data, a...)
}

func Error(format string, a ...interface{}) {
	now := getTime()
	var data string
	if checkOS() {
		data = fmt.Sprintf("[-] [%s] %s\n", now, format)
	} else {
		data = fmt.Sprintf("\x1b[31m[-] [%s] %s\x1b[0m\n", now, format)
	}
	_, _ = fmt.Fprintf(os.Stdout, data, a...)
}

func Warn(format string, a ...interface{}) {
	now := getTime()
	var data string
	if checkOS() {
		data = fmt.Sprintf("[!] [%s] %s\n", now, format)
	} else {
		data = fmt.Sprintf("\x1b[33m[!] [%s] %s\x1b[0m\n", now, format)
	}
	_, _ = fmt.Fprintf(os.Stdout, data, a...)
}

func getTime() string {
	currentTime := time.Now().Format("15:04:05")
	return currentTime
}

func checkOS() bool {
	sysType := runtime.GOOS
	if sysType == "windows" {
		return true
	}
	return false
}
