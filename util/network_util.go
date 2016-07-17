package util

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
)

const (
	HANDER     = "142857"
	HANDER_LEN = 6
	DATA_LEN   = 4
)

func Int64ToBytes(arg_value int64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(arg_value))
	return buf
}

func BytesToInt64(arg_bytes []byte) int64 {
	return int64(binary.BigEndian.Uint64(arg_bytes))
}

func EncryptPassword(arg_pwd string) string {
	hash := sha256.New()
	hash.Write([]byte(arg_pwd))
	return hex.EncodeToString(hash.Sum(nil))
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}
