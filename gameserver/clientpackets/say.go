package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

var (
	ALL                      int32 = 0
	AllChatRange             int32 = 1250  //Дальность белого чата
	Shout                    int32 = 1     //!
	ShoutChatRange           int32 = 10000 //Дальность желтого чата
	Tell                     int32 = 2     //\"
	PARTY                    int32 = 3     //#
	CLAN                     int32 = 4     //@
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

func Say(client *models.Client, data []byte) []byte {
	var packet = packets.NewReader(data)
	var say models.Say
	text := packet.ReadString()

	say.Text = text
	say.Type = packet.ReadInt32()

	switch say.Type {
	case ALL:
		//toBroad.B = serverpackets.CreatureSay(&say, me)
		//err := online.Char[me.CharId].Conn.Send(toBroad.GetB(), true) //to me
		//if err != nil {
		//	log.Println(err)
		//}
		//chars := models.GetAroundPlayersInRadius(me, AllChatRange)
		//for _, v := range chars {
		//	_ = online.Char[v].Conn.Send(toBroad.GetB(), true) //broad
		//}
	case Tell:
		//toBroad.B = serverpackets.CreatureSay(&say, me)
		//toTell := packet.ReadString()
	//	receiverExist := false
	//for _, v := range online.Char {
	//	if bytes.Equal(v.CharName.Bytes, []byte(toTell)) {
	//		receiverExist = true
	//		_ = v.Conn.Send(toBroad.GetB(), true)
	//
	//	}
	//}
	//	if receiverExist {
	//		_ = me.Conn.Send(toBroad.GetB(), true)
	//	} // else {
	//todo systemMSG not found
	// }
	case Shout:
		//toBroad.B = serverpackets.CreatureSay(&say, me)
		//err := online.Char[me.CharId].Conn.Send(toBroad.GetB(), true) //to me
		//if err != nil {
		//	log.Println(err)
		//}
		//chars := models.GetAroundPlayersInRadius(me, ShoutChatRange)
		//for _, v := range chars {
		//	_ = online.Char[v].Conn.Send(toBroad.GetB(), true) //broad
		//}
	}

	return []byte{}
}
