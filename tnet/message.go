package tnet

/*
 对链接内的数据封装， 消息
*/

type Message struct {
	// 消息id
	MsgId uint32
	// 消息长度
	MsgLen uint32
	// 消息具体数据
	MsgData []byte
}

func (m *Message) SetMsgData(data []byte) {
	m.MsgData = data
}

func (m *Message) GetMsgId() uint32 {
	return m.MsgId
}

func (m *Message) GetMsgData() []byte {
	return m.MsgData
}

func (m *Message) GetMsgLen() uint32 {
	return m.MsgLen
}
