package headler

import (
	"encoding/json"
	"fmt"
	"net"
	"servershell/bean"
	"servershell/bootstrap"
)

type MessLook struct {
}

func (m MessLook) IsInbound(shuju bootstrap.Pipemess) bool {
	return true
}
func (m MessLook) Dispose(client net.Conn, shuju bootstrap.Pipemess) bootstrap.Pipemess {
	key, ok := shuju.Value.(bean.Message)
	if !ok {
		return shuju
	}
	fmt.Println("MessLook:", key.MessId)
	switch key.MessId {
	case 1:
		user := bean.User{}
		json.Unmarshal(key.MessValue, &user)
		shuju.Value = user
		break
	default:
		shuju.Value = nil
		break
	}
	return shuju
}
