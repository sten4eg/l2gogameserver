package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"log"
)

var (
	ALL                      int32 = 0
	ALL_CHAT_RANGE           int32 = 1250 //Дальность белого чата
	SHOUT                    int32 = 1    //!
	TELL                     int32 = 2    //\"
	PARTY                    int32 = 3    //#
	CLAN                     int32 = 4    //@
	GM                       int32 = 5
	PETITION_PLAYER          int32 = 6 // used for petition
	PETITION_GM              int32 = 7 //* used for petition
	TRADE                    int32 = 8 //+
	ALLIANCE                 int32 = 9 //$
	ANNOUNCEMENT             int32 = 10
	PARTY_ROOM               int32 = 14
	COMMANDCHANNEL_ALL       int32 = 15 //`` (pink) команды лидера СС
	COMMANDCHANNEL_COMMANDER int32 = 16 //` (yellow) чат лидеров партий в СС
	HERO_VOICE               int32 = 17 //%
	CRITICAL_ANNOUNCEMENT    int32 = 18 //dark cyan
	UNKNOWN                  int32 = 19 //?
	BATTLEFIELD              int32 = 20 //^
)

func NewSay(data []byte, g *models.OnlineCharacters, me *models.Character) {
	var packet = packets.NewReader(data)
	var say models.Say
	text := packet.ReadString()

	say.Text = text
	say.Type = packet.ReadInt32()

	var info models.PacketByte
	info.B = serverpackets.NewCreatureSay(&say, me)
	err := g.Char[me.CharId].Conn.Send(info.GetB(), true)
	if err != nil {
		log.Println(err)
	}

	q := models.GetAroundPlayers(me, 200)

	for _, v := range q {
		g.Char[v].Conn.Send(info.GetB(), true)
	}
	//return &say
}
