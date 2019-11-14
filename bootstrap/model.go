package bootstrap

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
)

type ServerBootstrap struct {
	Post   string
	Mypipe Pipe
}

func (server *ServerBootstrap) SetPort(post string) {
	server.Post = post
}
func (server *ServerBootstrap) SetPipe(mypipe Pipe) {
	server.Mypipe = mypipe
}

func (server ServerBootstrap) Strat() {
	looker, err := net.Listen("tcp", server.Post)
	if err != nil {
		fmt.Println("无法绑定接口")
		return
	}
	fmt.Println("开始监听")
	for {
		client, err := looker.Accept()
		if err != nil {
			fmt.Println("创建客户端连接失败")
			break
		}
		pipeline := GetPipeline()
		ceshi := make([]Handler, 0)
		for _, handler := range server.Mypipe.Handlers {
			as := handler
			deepCopy(as, handler)
			ceshi = append(ceshi, as)
		}
		pipeline.Handlers = ceshi
		go pipeline.Start(client)
	}
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
