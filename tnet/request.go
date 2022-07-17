package tnet

import "github.com/itzmn/tin/tiface"

// Request 请求对链接和数据的封装
type Request struct {
	conn *Connection
	// 消息
	message tiface.IMessage
}

func (r *Request) GetMessage() tiface.IMessage {
	return r.message
}

func (r *Request) GetConnection() tiface.IConnection {
	return r.conn
}

func NewRequest(conn *Connection, message tiface.IMessage) *Request {

	return &Request{
		conn:    conn,
		message: message,
	}

}
