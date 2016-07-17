package main

import (
	"cst_im/controller"
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", ":6666" /* net.JoinHostPort(serverHost, strconv.Itoa(serverPort))*/)
	if err != nil {
		fmt.Printf("listening err %s:", err.Error())
		os.Exit(1)
	}
	fmt.Println("start accept")
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		fmt.Println(conn.RemoteAddr().String())
		my := controller.NewMyConnType(conn)
		listener := controller.SearchListener{Conn: my}
		my.Add(&listener)
		my.Read()
	}
}
