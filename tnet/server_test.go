package tnet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestServer_Start(t *testing.T) {

	name := "tinServer"
	ip := "127.0.0.1"
	port := 9999
	server := NewServer(name, ip, port)

	go server.Serve()
	time.Sleep(2 * time.Second)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}
	_, err = conn.Write([]byte{'a', 'b'})
	if err != nil {
		fmt.Println("write err:", err)
		return
	}

	buff := make([]byte, 512)
	cnt, err := conn.Read(buff)
	if err != nil {
		fmt.Println("read err:", err)
		return
	}
	fmt.Println("from server read data:", string(buff[:cnt]))
}
