package tnet

import (
	"fmt"
	"github.com/itzmn/tin/tiface"
	"net"
)

type Server struct {
	Name string
	IP   string
	Port int

	// server 处理业务的函数
	router *Router
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
	s.router.StartWorkerPool()
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
		connection := NewConnection(conn, cid, s)
		cid++
		connection.Start()
	}

}

func NewServer(name, ip string, port int) *Server {
	return &Server{
		Name:   name,
		IP:     ip,
		Port:   port,
		router: NewRouter(),
	}
}
func (s *Server) AddHandle(msgId uint32, handle tiface.IHandler) {
	s.router.AddHandler(msgId, handle)
}

func (s *Server) Stop() {
	// TODO 关闭系统资源
	fmt.Println("[tinServer]server stop")
}

func (s *Server) Serve() {
	fmt.Println("[tinServer]server serve")
	s.Start()

}
