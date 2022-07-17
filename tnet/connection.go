package tnet

import (
	"fmt"
	"github.com/itzmn/tin/tiface"
	"net"
)

type Connection struct {
	Server *Server
	// 链接id
	ConnectionId uint32
	// tcp socket链接
	Conn *net.TCPConn
	// 链接状态
	IsClose chan bool

	// 用户处理业务的函数
	handle tiface.IHandler
	// 用于读写分离的消息管道
	binaryDataChan chan []byte
}

// SendMsg 发送数据到客户端
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	dataPack := NewDataPack()
	msg := &Message{
		MsgId:   msgId,
		MsgLen:  uint32(len(data)),
		MsgData: data,
	}
	bytes, err := dataPack.Pack(msg)
	if err != nil {
		fmt.Printf("[tinServer]data pack err:%v", err)
		return err
	}
	c.binaryDataChan <- bytes

	return nil
}

func NewConnection(conn *net.TCPConn, id uint32, server *Server) *Connection {

	return &Connection{
		Conn:           conn,
		ConnectionId:   id,
		IsClose:        make(chan bool, 1),
		Server:         server,
		binaryDataChan: make(chan []byte),
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
	go c.startReader()
	// 启动从当前链接写数据的功能
	go c.startWriter()
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
func (c *Connection) startReader() {

	fmt.Println("[tinServer]Reader Goroutine is running...")
	defer fmt.Println("[tinServer]ConnectionId = ", c.ConnectionId, " Reader Goroutine is exit")
	defer c.Stop()

	for true {
		// 从链接读取数据，封装成为message
		message, err := ReadConnectionDataToMessage(c.GetTCPConnection())
		if err != nil {
			fmt.Println("[tinServer]connection readData to message err", err)
			return
		}
		// 2、封装request
		// 封装请求
		request := NewRequest(c, message)

		// 6、处理request
		// 调用当前链接处理数据的方法
		// 链接对请求进行处理，用户可以自定义处理逻辑
		go c.Server.router.SendRequestToTaskQueue(request)
		//go c.Server.router.DoHandle(request)

	}

}

// StartWriter 链接功能读写分离，主要从数据通道中获取数据 回写客户端
func (c *Connection) startWriter() {

	fmt.Println("[tinServer]Writer Goroutine is running...")
	defer fmt.Println("[tinServer]ConnectionId = ", c.ConnectionId, " Writer Goroutine is exit")

	for true {

		select {
		// 从数据channel中读取数据，回写客户端
		case data := <-c.binaryDataChan:
			_, err := c.Conn.Write(data)
			if err != nil {
				fmt.Printf("[tinServer]connId=%d, connection write data err:%v", c.GetConnId(), err)
				continue
			}
			// 如果connection已经关闭，不用再写了
		case <-c.IsClose:
			return

		}
	}
}

// ReadConnectionDataToMessage 从tcp链接按照TLV格式读取数据返回
func ReadConnectionDataToMessage(c net.Conn) (tiface.IMessage, error) {
	// 链接读取数据
	dataPack := NewDataPack()
	// 1、读取数据包头
	buffHeader := make([]byte, dataPack.GetPackHeaderLen())
	cnt, err := c.Read(buffHeader)
	if err != nil {
		fmt.Printf("[tinServer]connection read connection dataHeader err: %v\n", err)
		return nil, err
	}
	// 2、解包得到数据长度
	message, err := dataPack.UnPack(buffHeader)
	if err != nil {
		fmt.Printf("[tinServer]data unpack err:%v", err)
		return nil, err
	}
	// 3、读取对应长度的数据
	buff := make([]byte, message.GetMsgLen())
	cnt, err = c.Read(buff)
	if err != nil {
		fmt.Printf("[tinServer]connection read data err: %v\n", err)
		return nil, err
	}
	// 4、封装message
	message.SetMsgData(buff[:cnt])
	return message, nil
}
