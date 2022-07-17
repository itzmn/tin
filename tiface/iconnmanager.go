package tiface

/*
 链接管理抽象层
*/

type IConnManager interface {
	AddConn(connId uint32, connection IConnection)
	DelConn(connId uint32)
	Clear()
	GetConn(connId uint32) IConnection
	GetConnCnt() int
}
