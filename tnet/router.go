package tnet

import (
	"fmt"
	"github.com/itzmn/tin/tiface"
)

type Router struct {
	routerMap map[uint32]tiface.IHandler
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
