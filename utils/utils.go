package utils

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"unsafe"
)

func Int32ToBytes(data uint32) string {
	s := make([]byte, 4)
	binary.BigEndian.PutUint32(s, data)
	return hex.EncodeToString(s)
}
func Int16ToBytes(data uint16) string {
	s := make([]byte, 2)
	binary.BigEndian.PutUint16(s, data)
	return hex.EncodeToString(s)
}
func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
func Base64en(data []byte) string{
	encStr := base64.StdEncoding.EncodeToString([]byte(data))
	return encStr
}
func Base64de(enc_str string) string {
	decStr, _ := base64.StdEncoding.DecodeString(enc_str)
	return string(decStr)
}
