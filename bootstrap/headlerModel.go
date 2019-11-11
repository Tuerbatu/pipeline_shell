package bootstrap

import (
	"net"
)

type Handler struct {
	Mypipe *pipe
	Hfunc
}
type Hfunc interface {
	//必须实现方法,返回进出站类型
	IsInbound(shuju Pipemess) bool
	//必须实现方法，数据处理类
	Dispose(client net.Conn, shuju Pipemess) Pipemess
}

func (h *Handler) InitHandler(Pipe *pipe) {
	h.Mypipe = Pipe
}

func (h Handler) GetPipeline() pipe {
	return *h.Mypipe
}
