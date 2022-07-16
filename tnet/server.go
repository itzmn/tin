package tnet

import (
	"fmt"
	"net"
)

type Server struct {
	Name string
	IP   string
	Port int
}

func CallBackToClient(conn *net.TCPConn, buff []byte, cnt int) (err error) {
	fmt.Println("Connection Handle CallBackToClient...")
	if _, err = conn.Write(buff[:cnt]); err != nil {
		fmt.Println("connection write data to client err:", err)
	}
	return
}

func (s *Server) Start() {

	// 启动server监听
	listenAddr := fmt.Sprintf("%s:%v", s.IP, s.Port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", listenAddr)
	if err != nil {
		fmt.Println("[tinServer]resolve tcpAddr err:", err)
		return
	}
	listen, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		fmt.Println("[tinServer]listen ", listenAddr, " err, ", err)
		return
	}
	fmt.Println("[tinServer]server listen ", listenAddr, "success, listening...")
	var cid uint32
	cid = 0
	for true {
		// 等待连接
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("[tinServer]listen accept err, ", err)
			continue
		}
		fmt.Println("[tinServer]listen accept to handle")
		connection := NewConnection(conn, cid, CallBackToClient)
		cid++
		connection.Start()
	}

}

func NewServer(name, ip string, port int) *Server {
	return &Server{
		Name: name,
		IP:   ip,
		Port: port,
	}
}

func (s *Server) Stop() {
	// TODO 关闭系统资源
	fmt.Println("[tinServer]server stop")
}

func (s *Server) Serve() {
	fmt.Println("[tinServer]server serve")
	s.Start()

}
