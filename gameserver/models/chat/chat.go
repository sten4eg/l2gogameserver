package chat

const (
	All                      int32 = 0
	AllChatRange             int32 = 1250 //Дальность белого чата
	Shout                    int32 = 1    //!
	ShoutChatRange           int32 = 2000 //Дальность желтого чата
	Tell                     int32 = 2    //\"
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
	SpecialCommand           int32 = 21
)
