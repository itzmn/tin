package tnet

import (
	"fmt"
	"github.com/itzmn/tin/tiface"
	"net"
)

type Connection struct {

	// 链接id
	ConnectionId uint32
	// tcp socket链接
	Conn *net.TCPConn
	// 链接状态
	IsClose chan bool

	// 链接处理业务的函数
	handleFunc tiface.HandleFunc
}

func NewConnection(conn *net.TCPConn, id uint32, handleFunc tiface.HandleFunc) *Connection {

	return &Connection{
		Conn:         conn,
		ConnectionId: id,
		IsClose:      make(chan bool, 1),
		handleFunc:   handleFunc,
	}

}

func (c *Connection) GetConnId() uint32 {
	return c.ConnectionId
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) Start() (err error) {
	fmt.Println("[tinServer]Connection Start, Id = ", c.ConnectionId)
	// 启动从当前链接读取数据的功能
	go c.StartReader()
	// TODO 后续完善写数据的功能
	return
}
func (c *Connection) Stop() (err error) {
	fmt.Println("[tinServer]Connection Stop, Id =", c.ConnectionId)
	err = c.Conn.Close()
	c.IsClose <- true
	return
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// StartReader 读取数据
func (c *Connection) StartReader() {

	fmt.Println("[tinServer]Reader Goroutine is running...")
	defer fmt.Println("[tinServer]ConnectionId = ", c.ConnectionId, " Reader Goroutine is exit")
	defer c.Stop()

	for true {
		for true {
			buff := make([]byte, 512)
			cnt, err := c.GetTCPConnection().Read(buff)
			if err != nil {
				fmt.Printf("[tinServer]connection read data from client %s err: %v", c.GetRemoteAddr(), err)
				return
			}
			// 调用当前链接处理数据的方法
			fmt.Printf("[tinServer]connection read data:%v;  from client %s\n", string(buff[:cnt]), c.GetRemoteAddr())

			c.handleFunc(c.Conn, buff, cnt)

		}
	}

}
