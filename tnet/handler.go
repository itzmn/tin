package tnet

import (
	"github.com/itzmn/tin/tiface"
)

// BaseHandler 定义基础handler 为后续用户自定义作为父类
type BaseHandler struct{}

func (b *BaseHandler) PreHandle(request tiface.IRequest) {
}

func (b *BaseHandler) Handle(request tiface.IRequest) {
}

func (b *BaseHandler) PostHandle(request tiface.IRequest) {
}
