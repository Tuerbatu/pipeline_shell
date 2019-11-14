package bootstrap

import (
	"fmt"
	"net"
)

type Pipe struct {
	Handlers  []Handler
	Outmess   chan Pipemess
	IsSurvive bool
}
type Pipemess struct {
	IsInbound bool
	Value     interface{}
}

func (p *Pipe) AddHeadler(h Handler) {
	p.Handlers = append(p.Handlers, h)
}

func GetPipeline() Pipe {
	return Pipe{make([]Handler, 0), make(chan Pipemess, 5), true}
}

func (p *Pipe) doWrite(client net.Conn) {
	for {
		shuju, err := <-p.Outmess
		if err || !p.IsSurvive {
			return
		}
		ok := false
		for i := len(p.Handlers) - 1; i >= 0; i-- {
			shuju, ok = manage(client, p.Handlers[i], p, shuju)
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
func (p Pipe) Write(data []byte) {
	key := Pipemess{false, data}
	p.Outmess <- key
}

func (p *Pipe) Close() {
	p.IsSurvive = false
}

func (p *Pipe) Start(client net.Conn) {
	fmt.Println("连接成功,Handler数量:", len(p.Handlers))
	go p.doWrite(client)
	for {
		if !p.IsSurvive {
			return
		}
		shuju := Pipemess{true, nil}
		ok := false
		for _, handler := range p.Handlers {
			shuju, ok = manage(client, handler, p, shuju)
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
func manage(client net.Conn, key Handler, pipe *Pipe, shuju Pipemess) (Pipemess, bool) {
	if key.IsInbound(shuju) != shuju.IsInbound {
		return shuju, false
	}
	jieguo := key.Dispose(client, pipe, shuju)
	return jieguo, true
}
