package tiface

/*
 针对数据包进行封装，应对拆包、黏包问题
*/

type IDataPack interface {

	// GetPackHeaderLen 获取包头长度
	GetPackHeaderLen() uint32
	// Pack 将message封装为数据包
	Pack(message IMessage) ([]byte, error)
	// UnPack 将tcp链接数据，转换成message
	UnPack(data []byte) (IMessage, error)
}
