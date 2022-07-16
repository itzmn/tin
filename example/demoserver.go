package main

import (
	"fmt"
	"github.com/itzmn/tin/tiface"
	"github.com/itzmn/tin/tnet"
)

type PingHandle struct {
	tnet.BaseHandler
}

func (p *PingHandle) PreHandle(request tiface.IRequest) {
	fmt.Println("PingHandle preHandle, connectionId=", request.GetConnection().GetConnId())
}

func (p *PingHandle) Handle(request tiface.IRequest) {
	fmt.Println("PingHandle Handle request, connectionId=", request.GetConnection().GetConnId())
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("PingHandle return Ping Ping Ping"))
	if err != nil {
		fmt.Println("PingHandle handle err:", err)
	}
}

func (p *PingHandle) PostHandle(request tiface.IRequest) {
	fmt.Println("PingHandle postHandle request, connectionId=", request.GetConnection().GetConnId())
}

func main() {
	name := "tinServer"
	ip := "127.0.0.1"
	port := 9999
	server := tnet.NewServer(name, ip, port)
	server.AddHandle(&PingHandle{})

	server.Serve()
}
