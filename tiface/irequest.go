package tiface

type IRequest interface {
	// GetConnection 获取请求链接
	GetConnection() IConnection
	// GetData 获取请求数据
	GetData() []byte
}
