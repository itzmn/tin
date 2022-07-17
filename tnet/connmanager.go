package tnet

import (
	"fmt"
	"github.com/itzmn/tin/tiface"
	"sync"
)

type ConnManager struct {
	connMap map[uint32]tiface.IConnection
	lock    sync.RWMutex
}

func (c *ConnManager) AddConn(connId uint32, connection tiface.IConnection) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if _, ok := c.connMap[connId]; ok {
		fmt.Printf("connId = %d exist on connMap", connId)
		return
	}
	c.connMap[connId] = connection
}

func (c *ConnManager) DelConn(connId uint32) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if _, ok := c.connMap[connId]; !ok {
		fmt.Printf("connId = %d no exist on connMap", connId)
		return
	}
	delete(c.connMap, connId)
}

func (c *ConnManager) Clear() {
	c.lock.Lock()
	defer c.lock.Unlock()
	for id, _ := range c.connMap {
		delete(c.connMap, id)
	}

}

func (c *ConnManager) GetConn(connId uint32) tiface.IConnection {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if _, ok := c.connMap[connId]; !ok {
		fmt.Printf("connId = %d no exist on connMap", connId)
		return nil
	}
	return c.connMap[connId]
}

func (c *ConnManager) GetConnCnt() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return len(c.connMap)
}

func NewConnManager() *ConnManager {

	return &ConnManager{
		connMap: make(map[uint32]tiface.IConnection),
	}
}
