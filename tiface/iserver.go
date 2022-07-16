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
}
