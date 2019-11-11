package bootstrap

import (
	"fmt"
	"net"
)

type pipe struct {
	Handlers []Handler
	Outmess  chan Pipemess
}
type Pipemess struct {
	IsInbound bool
	Value     interface{}
}

func (p *pipe) AddHeadler(h Handler) {
	h.InitHandler(p)
	p.Handlers = append(p.Handlers, h)
}

func GetPipeline() pipe {
	return pipe{make([]Handler, 0), make(chan Pipemess, 5)}
}

func (p pipe) doWrite(client net.Conn) {
	for {
		shuju := <-p.Outmess
		ok := false
		for i := len(p.Handlers) - 1; i >= 0; i-- {
			shuju, ok = manage(client, p.Handlers[i], shuju)
			if !ok {
				continue
			}
			fmt.Println(shuju.Value)
			if shuju.Value == nil {
				fmt.Println("没有数据停止深入")
				break
			}
		}
	}
}
func (p pipe) Write(data []byte) {
	key := Pipemess{false, data}
	p.Outmess <- key
}

func (p pipe) Start(client net.Conn) {
	fmt.Println("连接成功,Handler数量:", len(p.Handlers))
	go p.doWrite(client)
	for {
		shuju := Pipemess{true, nil}
		ok := false
		for _, handler := range p.Handlers {
			shuju, ok = manage(client, handler, shuju)
			if !ok {
				continue
			}
			fmt.Println(shuju.Value)
			if shuju.Value == nil {
				fmt.Println("没有数据停止深入")
				break
			}
		}
	}
}
func manage(client net.Conn, key Hfunc, shuju Pipemess) (Pipemess, bool) {
	if key.IsInbound(shuju) != shuju.IsInbound {
		return shuju, false
	}
	jieguo := key.Dispose(client, shuju)
	return jieguo, true
}
