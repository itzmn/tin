package tnet

import (
	"bytes"
	"encoding/binary"
	"github.com/itzmn/tin/tiface"
)

// header  |msgId  |data
// 4byte   |4byte  | len
// dataLen | msgId | 数据

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (d *DataPack) GetPackHeaderLen() uint32 {
	return 8
}

// Pack 将messgae通过二进制方式序列化为 TLV格式数据包
func (d *DataPack) Pack(message tiface.IMessage) ([]byte, error) {
	// 创建存储二进制数据的数组
	buffer := bytes.NewBuffer([]byte{})
	// 将dataLen 写入二进制
	if err := binary.Write(buffer, binary.LittleEndian, message.GetMsgLen()); err != nil {
		return nil, err
	}
	// 将MsgId写入
	if err := binary.Write(buffer, binary.LittleEndian, message.GetMsgId()); err != nil {
		return nil, err
	}
	// 将Msg数据写入
	if err := binary.Write(buffer, binary.LittleEndian, message.GetMsgData()); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// UnPack 将字节数组按照TLV格式反序列化为message
func (d *DataPack) UnPack(data []byte) (tiface.IMessage, error) {
	// 创建一个从二进制读取的buffer
	reader := bytes.NewReader(data)
	// 只解压读取header
	msg := &Message{}
	// 读取msgLen
	if err := binary.Read(reader, binary.LittleEndian, &msg.MsgLen); err != nil {
		return nil, err
	}
	// 读取msgId
	if err := binary.Read(reader, binary.LittleEndian, &msg.MsgId); err != nil {
		return nil, err
	}
	return msg, nil
}
