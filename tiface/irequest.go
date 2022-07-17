package tiface

type IRequest interface {
	// GetConnection 获取请求链接
	GetConnection() IConnection
	// GetMessage 获取请求数据
	GetMessage() IMessage
}
