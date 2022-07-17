package main

import (
	"fmt"
	"github.com/itzmn/tin/config"
	"github.com/itzmn/tin/tiface"
	"github.com/itzmn/tin/tnet"
)

type PingHandle struct {
	tnet.BaseHandler
}

func (p *PingHandle) Handle(request tiface.IRequest) {
	fmt.Printf("PingHandle Handle request,connectionId=%d, msgId=%d, msgData:%v\n",
		request.GetConnection().GetConnId(), request.GetMessage().GetMsgId(), string(request.GetMessage().GetMsgData()))
	err := request.GetConnection().SendMsg(5, []byte("PingHandle return Ping Ping Ping"))
	if err != nil {
		fmt.Println("PingHandle handle err:", err)
	}
}

type PongHandle struct {
	tnet.BaseHandler
}

func (p *PongHandle) Handle(request tiface.IRequest) {
	fmt.Printf("PongHandle Handle request,connectionId=%d, msgId=%d, msgData:%v\n",
		request.GetConnection().GetConnId(), request.GetMessage().GetMsgId(), string(request.GetMessage().GetMsgData()))
	err := request.GetConnection().SendMsg(5, []byte("PongHandle return Pong Pong Pong"))
	if err != nil {
		fmt.Println("PongHandle handle err:", err)
	}
}

func main() {

	name := config.GConfig.ServerName
	ip := config.GConfig.IP
	port := config.GConfig.Port
	server := tnet.NewServer(name, ip, port)
	server.AddHandle(1, &PingHandle{})
	server.AddHandle(2, &PongHandle{})

	server.SetOnConnStart(func(connection tiface.IConnection) {
		fmt.Println("this is hook func, onConnStart, connId=", connection.GetConnId())
	})
	server.Serve()
}
