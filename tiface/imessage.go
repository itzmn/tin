package tiface

/*
 对链接内的数据封装， 消息
*/

type IMessage interface {
	GetMsgId() uint32
	GetMsgData() []byte
	GetMsgLen() uint32
	SetMsgData(data []byte)
}
