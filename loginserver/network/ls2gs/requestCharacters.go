package ls2gs

import (
	"l2gogameserver/loginserver/network/gs2ls"
	"l2gogameserver/packets"
)

type loginServerInterfaceRC interface {
	Send(buffer *packets.Buffer)
}

func RequestCharacters(data []byte, ls loginServerInterfaceRC) {
	reader := packets.NewReader(data)
	login := reader.ReadString()

	ls.Send(gs2ls.ReplyCharacters(login))
}
