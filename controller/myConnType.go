package controller

import (
	pb "cst_im/model"
	"cst_im/util"
	"fmt"
	"net"

	"github.com/golang/protobuf/proto"
)

type MyConnType struct {
	conn         net.Conn
	isConnected  bool
	MsgChannel   chan pb.Msg
	listenerList []iListener
}

func NewMyConnType(arg_conn net.Conn) *MyConnType {
	myConnType := &MyConnType{
		conn:        arg_conn,
		isConnected: true,
	}
	myConnType.MsgChannel = make(chan pb.Msg)
	myConnType.listenerList = make([]iListener, 0)
	go func() {
		for myConnType.isConnected {
			msg := <-myConnType.MsgChannel
			fmt.Println(msg)
			for i := 0; i < len(myConnType.listenerList); i++ {
				myConnType.listenerList[i].onProcess(msg)
			}
		}
	}()
	return myConnType
}
func (this *MyConnType) Add(l iListener) {
	this.listenerList = append(this.listenerList, l)
}

func (this *MyConnType) Read() {
	go func() {
		buf := make([]byte, 1024)
		tempBuf := make([]byte, 0)
		for {
			n, err := this.conn.Read(buf)
			if err != nil {
				fmt.Println(err)
				this.isConnected = false
				break
			}
			tempBuf, err = this.unpack(append(tempBuf, buf[:n]...))
		}
	}()

}

func (this *MyConnType) Write(msg *pb.Msg) {
	buf, err := this.pack(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("data_length: ", len(buf))
	Len := 4 * 1024     //每次传输的数据为4k
	dataLen := len(buf) //数据长度
	if dataLen <= Len { //如果数据小于4K
		this.conn.Write(buf)
		return
	}
	i := 0
	for ; dataLen-i*Len >= Len; i++ {
		this.conn.Write(buf[i*Len : i*Len+Len])
	}
	restOfLen := dataLen - i*Len
	if restOfLen <= 0 {
		return
	}
	this.conn.Write(buf[i*Len:])
}

func (this *MyConnType) unpack(arg_buf []byte) ([]byte, error) {
	totalLen := len(arg_buf)
	var i int
	for i = 0; i < totalLen; i++ {
		if totalLen < i+util.HANDER_LEN+util.DATA_LEN {
			break
		}
		if string(arg_buf[i:i+util.HANDER_LEN]) == util.HANDER {
			dataLen := util.BytesToInt(arg_buf[i+util.HANDER_LEN : i+util.HANDER_LEN+util.DATA_LEN])
			if totalLen < dataLen+util.HANDER_LEN+util.DATA_LEN {
				break
			}

			data := arg_buf[i+util.HANDER_LEN+util.DATA_LEN : i+util.HANDER_LEN+util.DATA_LEN+dataLen]
			msg := &pb.Msg{}
			err := proto.Unmarshal(data, msg)
			if err != nil {
				fmt.Println(err)
				break
			}
			this.MsgChannel <- *msg
			i = i + util.HANDER_LEN + util.DATA_LEN + dataLen - 1
		}
	}
	if i == totalLen {
		return make([]byte, 0), nil
	}
	return arg_buf[i:], nil
}

func (this *MyConnType) pack(msg *pb.Msg) ([]byte, error) {
	dataBuf, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	dataLen := len(dataBuf)
	dataLenBytes := util.IntToBytes(dataLen)
	buf := append([]byte(util.HANDER), dataLenBytes...)
	return append(buf, dataBuf...), nil
}
