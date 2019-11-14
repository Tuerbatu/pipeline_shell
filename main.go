package main

import (
	"github.com/Tuerbatu/pipeline_shell/bootstrap"
	"github.com/Tuerbatu/pipeline_shell/headler"
)

func main() {
	pipeline := bootstrap.GetPipeline()
	pipeline.AddHeadler(headler.ByteToMess{
		Buffer:   make([]byte, 200),
		Readline: make([]byte, 200),
	})
	pipeline.AddHeadler(headler.MessLook{})
	pipeline.AddHeadler(headler.Login{})
	server := bootstrap.ServerBootstrap{":8080", pipeline}
	server.Strat()
}
