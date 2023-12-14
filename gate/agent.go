package gate

import (
	"net"
)

type Agent interface {
	WriteMsg(msg interface{})
	SendMsg(opCode uint32, msg interface{})
	SendBuff(opCode uint32, data [][]byte)
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
}
