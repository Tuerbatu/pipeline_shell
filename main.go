package main

import (
	"servershell/bootstrap"
	"servershell/headler"
)

func main() {
	pipeline := bootstrap.GetPipeline()
	pipeline.AddHeadler(bootstrap.Handler{
		Hfunc: &headler.ByteToMess{
			Buffer:   make([]byte, 200),
			Readline: make([]byte, 200),
		}})
	pipeline.AddHeadler(bootstrap.Handler{
		Hfunc: &headler.MessLook{}})
	pipeline.AddHeadler(bootstrap.Handler{
		Hfunc: &headler.Login{}})
	server := bootstrap.ServerBootstrap{":8080", pipeline}
	server.Strat()
}
