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

func (s *Server) Start() {

	// 启动server监听
	listenAddr := fmt.Sprintf("%s:%v", s.IP, s.Port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", listenAddr)
	if err != nil {
		fmt.Println("resolve tcpAddr err:", err)
		return
	}
	listen, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		fmt.Println("listen ", listenAddr, " err, ", err)
		return
	}
	fmt.Println("server listen ", listenAddr, "success, listening...")
	for true {
		// 等待连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen accept err, ", err)
			continue
		}
		fmt.Println("listen accept to handle")
		go func() {

			for true {
				buff := make([]byte, 512)
				cnt, err := conn.Read(buff)
				if err != nil {
					fmt.Println("conn read data err: ", err)
					return
				}
				//res := string(buff)
				fmt.Printf("recive buf:%v, cnt:%d\n", string(buff[:cnt]), cnt)
				if _, err := conn.Write(buff[:cnt]); err != nil {
					fmt.Println("rewrite err:", err)
				}
			}

		}()
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
	fmt.Println("[tin] server stop")
}

func (s *Server) Serve() {
	fmt.Println("[tin] server serve")
	s.Start()

}
