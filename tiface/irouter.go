package tiface

/*
 路由模块，针对不同的请求做不同的业务处理
*/

type IRouter interface {
	// AddHandler 给路由增加handle
	AddHandler(msgId uint32, handler IHandler)
	// DoHandle 执行业务的方法
	DoHandle(request IRequest)
	// StartWorkerPool 开启线程池处理任务
	StartWorkerPool()
	// SendRequestToTaskQueue 将消息发送给消息队列
	SendRequestToTaskQueue(request IRequest)
}
