package gate

import (
	"net"
)

type Agent interface {
	WriteMsg(msg interface{})
	SendMsg(opCode uint32, msg interface{})
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
}
