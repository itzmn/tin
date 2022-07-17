package main

import (
	"fmt"
	"github.com/itzmn/tin/config"
	"github.com/itzmn/tin/tnet"
	"math/rand"
	"net"
	"time"
)

func main() {
	ip := config.GConfig.IP
	port := config.GConfig.Port
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}
	for true {
		// 将数据打包发送给server端
		dataPack := tnet.NewDataPack()
		msg := &tnet.Message{
			MsgId:   uint32(rand.Int()%2 + 1),
			MsgLen:  2,
			MsgData: []byte{'a', 'b'},
		}
		bytes, err := dataPack.Pack(msg)
		_, err = conn.Write(bytes)
		if err != nil {
			fmt.Println("write err:", err)
			return
		}

		// 读取server 回写的数据
		message, err := tnet.ReadConnectionDataToMessage(conn)
		//buff := make([]byte, 512)
		//cnt, err := conn.Read(buff)
		if err != nil {
			fmt.Println("read err:", err)
			return
		}
		fmt.Println("from server read data:", string(message.GetMsgData()))
		time.Sleep(3 * time.Second)
	}
	fmt.Println("client end...")

}
