package tnet

import (
	"fmt"
	"github.com/itzmn/tin/config"
	"github.com/itzmn/tin/tiface"
)

type Router struct {
	routerMap map[uint32]tiface.IHandler
	// 消息队列，用于异步处理任务
	TaskQueue []chan tiface.IRequest
	// 任务处理携程个数
	WorkPoolSize uint32
}

func NewRouter() *Router {
	return &Router{
		routerMap:    make(map[uint32]tiface.IHandler),
		WorkPoolSize: uint32(config.GConfig.WorkerPoolSize),
		TaskQueue:    make([]chan tiface.IRequest, config.GConfig.WorkerPoolSize),
	}
}

// SendRequestToTaskQueue 发送请求给消息队列，由后台任务进行处理
func (r *Router) SendRequestToTaskQueue(request tiface.IRequest) {

	routeId := request.GetConnection().GetConnId() % r.WorkPoolSize
	r.TaskQueue[routeId] <- request
	fmt.Printf("send connectId=%d request to workerId=%d taskQueue\n", request.GetConnection().GetConnId(), routeId)

}

func (r *Router) StartWorkerPool() {

	for i := 0; i < int(r.WorkPoolSize); i++ {
		r.TaskQueue[i] = make(chan tiface.IRequest, r.WorkPoolSize)
		go r.startOneWorker(i)
	}
}

func (r *Router) startOneWorker(workerId int) {
	fmt.Println("start worker, id=", workerId)
	requests := r.TaskQueue[workerId]
	for true {
		select {
		case request := <-requests:
			r.DoHandle(request)
		}
	}
}

func (r *Router) AddHandler(msgId uint32, handler tiface.IHandler) {
	if _, ok := r.routerMap[msgId]; ok {
		fmt.Printf("router msgId=%d exist in routerMap\n", msgId)
		return
	}
	fmt.Printf("add router msgId=%d to routerMap\n", msgId)
	r.routerMap[msgId] = handler
}

// DoHandle 根据请求类型在路由中寻找处理函数处理
func (r *Router) DoHandle(request tiface.IRequest) {
	msgId := request.GetMessage().GetMsgId()
	handle, ok := r.routerMap[msgId]
	if !ok {
		fmt.Printf("router msgId=%d no exist in routerMap\n", msgId)
		return
	}
	handle.PreHandle(request)
	handle.Handle(request)
	handle.PostHandle(request)
}
