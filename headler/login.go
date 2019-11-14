package headler

import (
	"fmt"
	"net"

	"github.com/Tuerbatu/pipeline_shell/bean"
	"github.com/Tuerbatu/pipeline_shell/bootstrap"
)

type Login struct {
}

func (l Login) IsInbound(shuju bootstrap.Pipemess) bool {
	return true
}
func (l Login) Dispose(client net.Conn, mypipe *bootstrap.Pipe, shuju bootstrap.Pipemess) bootstrap.Pipemess {
	key, ok := shuju.Value.(bean.User)
	if !ok {
		return shuju
	}
	fmt.Println("MessLook:", key.UserName)
	if key.UserName == "tuerbatu" && key.Password == "123456" {
		fmt.Println("登录成功")
	}
	shuju.Value = nil
	return shuju
}
