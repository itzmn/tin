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

	// 用户处理业务的函数
	handle tiface.IHandler
}

func NewConnection(conn *net.TCPConn, id uint32, handle tiface.IHandler) *Connection {

	return &Connection{
		Conn:         conn,
		ConnectionId: id,
		IsClose:      make(chan bool, 1),
		handle:       handle,
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
		// 链接读取数据
		dataPack := NewDataPack()
		// 1、读取数据包头
		buffHeader := make([]byte, dataPack.GetPackHeaderLen())
		cnt, err := c.GetTCPConnection().Read(buffHeader)
		if err != nil {
			fmt.Printf("[tinServer]connection read connection dataHeader from client %s err: %v\n", c.GetRemoteAddr(), err)
			return
		}
		// 2、解包得到数据长度
		message, err := dataPack.UnPack(buffHeader)
		if err != nil {
			fmt.Printf("[tinServer]data unpack err:%v", err)
			return
		}
		// 3、读取对应长度的数据
		buff := make([]byte, message.GetMsgLen())
		cnt, err = c.GetTCPConnection().Read(buff)
		if err != nil {
			fmt.Printf("[tinServer]connection read data from client %s err: %v\n", c.GetRemoteAddr(), err)
			return
		}
		fmt.Printf("[tinServer]connection read data:%v;  from client %s\n", string(buff[:cnt]), c.GetRemoteAddr())
		// 4、封装message
		message.SetMsgData(buff[:cnt])
		// 5、封装request
		// 封装请求
		request := NewRequest(c, buff, message)
		// 6、处理request
		// 调用当前链接处理数据的方法
		// 链接对请求进行处理，用户可以自定义处理逻辑
		go func(request tiface.IRequest) {
			c.handle.PreHandle(request)
			c.handle.Handle(request)
			c.handle.PostHandle(request)
		}(request)

	}

}
