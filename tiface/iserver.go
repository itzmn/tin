package tiface

/*
 server 接口
*/

type IServer interface {

	// Start 启动server
	Start()
	// Stop 关闭server
	Stop()
	// Serve 执行server逻辑
	Serve()

	// AddHandle 增加处理业务的逻辑
	AddHandle(msgId uint32, handle IHandler)

	SetOnConnStart(func(connection IConnection))
	SetOnConnStop(func(connection IConnection))
}
