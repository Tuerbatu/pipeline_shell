package bootstrap

import (
	"net"
)

type Handler interface {
	//必须实现方法,返回进出站类型
	IsInbound(shuju Pipemess) bool
	//必须实现方法，数据处理类
	Dispose(client net.Conn, mypipe *Pipe, shuju Pipemess) Pipemess
}
