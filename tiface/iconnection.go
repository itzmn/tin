package tiface

import "net"

type IConnection interface {
	GetConnId() uint32

	GetTCPConnection() *net.TCPConn

	Stop() (err error)
	Start() (err error)

	GetRemoteAddr() net.Addr
}

// HandleFunc 定义一个处理业务的函数类型
type HandleFunc func(conn *net.TCPConn, msg []byte, cnt int) (err error)
