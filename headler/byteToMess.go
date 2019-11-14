package headler

import (
	"encoding/json"
	"net"

	"github.com/Tuerbatu/pipeline_shell/bean"
	"github.com/Tuerbatu/pipeline_shell/bootstrap"
)

type ByteToMess struct {
	Buffer   []byte
	Readline []byte
}

func (b ByteToMess) IsInbound(shuju bootstrap.Pipemess) bool {
	return true
}
func (b ByteToMess) Dispose(client net.Conn, mypipe *bootstrap.Pipe, shuju bootstrap.Pipemess) bootstrap.Pipemess {
	i, _ := client.Read(b.Readline)
	key := []byte{}
	switch b.Readline[0] {
	case 0:
		key = b.Readline[1:i]
		break
	case 1:
		b.Buffer = make([]byte, len(b.Readline))
	case 2:
		b.Buffer = append(b.Buffer, b.Readline[1:i]...)
		key = nil
		break
	case 3:
		b.Buffer = append(b.Buffer, b.Readline[1:i]...)
		key = b.Buffer
		break
	}
	if key == nil {
		shuju.Value = nil
	} else {
		jieguo := bean.Message{}
		json.Unmarshal(key, &jieguo)
		shuju.Value = jieguo
	}
	return shuju
}
