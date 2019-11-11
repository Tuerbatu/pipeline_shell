package headler

import (
	"fmt"
	"net"
	"servershell/bean"
	"servershell/bootstrap"
)

type Login struct {
}

func (l Login) IsInbound(shuju bootstrap.Pipemess) bool {
	return true
}
func (l Login) Dispose(client net.Conn, shuju bootstrap.Pipemess) bootstrap.Pipemess {
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
