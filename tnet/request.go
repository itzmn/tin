package tnet

import "github.com/itzmn/tin/tiface"

// Request 请求对链接和数据的封装
type Request struct {
	conn *Connection
	data []byte
}

func (r *Request) GetConnection() tiface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}

func NewRequest(conn *Connection, data []byte) *Request {

	return &Request{
		conn: conn,
		data: data,
	}

}
