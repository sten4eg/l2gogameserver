package models

import (
	"context"
	"fmt"
	"l2gogameserver/data"
	"l2gogameserver/data/logger"
	"l2gogameserver/db"
	"l2gogameserver/gameserver/dto"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/clientStates"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/models/race"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
	"net"

	"sync"
	"time"
)

type (
	Character struct {
		Login         string
		ObjectId      int32
		Level         int32
		MaxHp         int32
		CurHp         int32
		MaxMp         int32
		CurMp         int32
		Face          int32
		HairStyle     int32
		HairColor     int32
		Sex           int32
		Coordinates   *Coordinates
		Heading       int32
		Exp           int32
		Sp            int32
		Karma         int32
		PvpKills      int32
		PkKills       int32
		ClanId        int32
		Race          race.Race
		ClassId       int32
		BaseClass     int32
		Title         string
		OnlineTime    int32
		Nobless       int32
		Vitality      int32
		CharName      string
		CurrentRegion *WorldRegion
		Conn          *ClientCtx
		SockConn      *net.TCPConn
		AttackEndTime int64
		// Paperdoll - массив всех слотов которые можно одеть
		Paperdoll       [26]MyItem
		Stats           StaticData
		pvpFlag         bool
		ShortCut        map[int32]dto.ShortCutDTO
		ActiveSoulShots []int32
		IsDead          bool
		IsFakeDeath     bool
		// Skills todo: проверить слайс или мапа лучше для скилов
		Skills                  map[int]Skill
		IsCastingNow            bool
		SkillQueue              chan SkillHolder
		CurrentSkill            *SkillHolder // todo А может быть без * попробовать?
		Inventory               Inventory
		CursedWeaponEquippedId  int
		BonusStats              []items.ItemBonusStat
		ChannelUpdateShadowItem chan IUP
		InGame                  bool
		Target                  int32
		Macros                  []Macro
		CharInfoTo              chan []int32
		DeleteObjectTo          chan []int32
		NpcInfo                 chan []interfaces.Npcer
		IsMoving                bool
		Sit                     bool
		FirstEnterGame          bool
		ActiveRequester         interfaces.CharacterI
		RequestExpireTime       int64
		ActiveTradeList         *TradeList
		TradeRefusal            bool
		ActiveEnchantItemId     int32
	}
	SkillHolder struct {
		Skill        Skill
		CtrlPressed  bool
		ShiftPressed bool
	}
	Coordinates struct {
		mu sync.Mutex
		X  int32
		Y  int32
		Z  int32
	}
	ToSendInfo struct {
		To   []int32
		Info utils.PacketByte
	}

	IUP struct {
		ObjId      int32
		UpdateType int16
	}
)

const (
	RequestTimeout = time.Second * 10
	IdNone         = -1
)

func GetNewCharacterModel() *Character {
	character := new(Character)
	sk := make(map[int]Skill)
	character.Skills = sk
	character.ChannelUpdateShadowItem = make(chan IUP, 10)
	character.InGame = false
	character.ActiveEnchantItemId = IdNone
	return character
}

// SetSitStandPose Меняет положение персонажа от сидячего к стоячему и на оборот
// Возращает значение нового положения
func (c *Character) SetSitStandPose() int32 {
	if !c.Sit {
		c.Sit = true
		return 0
	}
	c.Sit = false
	return 1
}

func (c *Character) ListenSkillQueue() {
	for {
		select {
		case res := <-c.SkillQueue:
			fmt.Println("SKILL V QUEUE")
			fmt.Println(res)
		}
	}
}

func (c *Character) SetSkillToQueue(skill Skill, ctrlPressed, shiftPressed bool) {
	s := SkillHolder{
		Skill:        skill,
		CtrlPressed:  ctrlPressed,
		ShiftPressed: shiftPressed,
	}
	c.SkillQueue <- s
}

// IsActiveWeapon есть ли у персонажа оружие в руках
func (c *Character) IsActiveWeapon() bool {
	x := c.Paperdoll[PAPERDOLL_RHAND]
	//todo Еще есть кастеты
	return x.ObjectId != 0
	//todo ?
}

// GetPercentFromCurrentLevel получить % опыта на текущем уровне
func (c *Character) GetPercentFromCurrentLevel(exp, level int32) float64 {
	expPerLevel, expPerLevel2 := data.GetExpData(level)
	return float64(int64(exp)-expPerLevel) / float64(expPerLevel2-expPerLevel)
}

// Load загрузка персонажа
func (c *Character) Load() {
	c.InGame = true
	c.ShortCut = RestoreMe(c.ObjectId, c.ClassId)
	c.LoadSkills()
	c.SkillQueue = make(chan SkillHolder)
	c.Inventory.Items = GetMyItems(c.ObjectId)
	c.Paperdoll = RestoreVisibleInventory(c.ObjectId)
	c.LoadCharactersMacros()
	for _, v := range &c.Paperdoll {
		if v.ObjectId != 0 {
			c.AddBonusStat(v.BonusStats)
		}
	}
	c.Stats = AllStats[int(c.ClassId)].StaticData //todo а для чего BaseClass ??

	reg := GetRegion(c.Coordinates.X, c.Coordinates.Y, c.Coordinates.Z)
	c.CharInfoTo = make(chan []int32, 2)
	c.DeleteObjectTo = make(chan []int32, 2)
	c.NpcInfo = make(chan []interfaces.Npcer, 2)
	c.setWorldRegion(reg)

	reg.AddVisibleChar(c)

	go c.Shadow()
	go c.ListenSkillQueue()
	go c.checkRegion()

}

func (c *Character) Shadow() {
	for {
		for i := range c.Inventory.Items {
			v := &c.Inventory.Items[i]
			if v.Item.Durability > 0 && v.Location == PaperdollLoc {
				var iup IUP
				iup.ObjId = v.ObjectId
				switch c.Inventory.Items[i].Mana {

				case 0:
					iup.UpdateType = UpdateTypeRemove
					c.ChannelUpdateShadowItem <- iup
					DeleteItem(v, c)
				default:
					c.Inventory.Items[i].Mana -= 1
					iup.UpdateType = UpdateTypeModify
					c.ChannelUpdateShadowItem <- iup
				}

			}
		}

		time.Sleep(time.Second)
	}

}

func (c *Character) checkSoulShot() {
	if len(c.ActiveSoulShots) == 0 {
		return
	}
}

func (c *Character) IsCursedWeaponEquipped() bool {
	return c.CursedWeaponEquippedId != 0
}

func (c *Character) AddBonusStat(s []items.ItemBonusStat) {
	c.BonusStats = append(c.BonusStats, s...)
}

func (c *Character) RemoveBonusStat(s []items.ItemBonusStat) {
	//for i,v := range c.BonusStats {
	//	for _,vv := range s {
	//		if v == vv {
	//			c.BonusStats[i] = c.BonusStats[len(c.BonusStats)-1] //todo переделать на безопасный вариант ) или еще что нить придумать
	//			c.BonusStats = c.BonusStats[:len(c.BonusStats)-1]
	//		}
	//	}
	//
	//}

	news := make([]items.ItemBonusStat, 0, len(c.BonusStats))
	for _, v := range c.BonusStats {
		flag := false
		for _, vv := range s {
			if v == vv {
				flag = true
				break
			}
		}
		if !flag {
			news = append(news, v)
		}
	}
	c.BonusStats = news
}

func (c *Character) GetPDef() int32 {
	var base float64
	if c.Paperdoll[PAPERDOLL_FEET].ObjectId == 0 {
		base = float64(c.Stats.BasePDef.Feet)
	}
	if c.Paperdoll[PAPERDOLL_CHEST].ObjectId == 0 {
		base += float64(c.Stats.BasePDef.Chest)
	}
	if c.Paperdoll[PAPERDOLL_CLOAK].ObjectId == 0 {
		base += float64(c.Stats.BasePDef.Cloak)
	}
	if c.Paperdoll[PAPERDOLL_HEAD].ObjectId == 0 {
		base += float64(c.Stats.BasePDef.Head)
	}
	if c.Paperdoll[PAPERDOLL_GLOVES].ObjectId == 0 {
		base += float64(c.Stats.BasePDef.Gloves)
	}
	if c.Paperdoll[PAPERDOLL_LEGS].ObjectId == 0 {
		base += float64(c.Stats.BasePDef.Legs)
	}
	if c.Paperdoll[PAPERDOLL_UNDER].ObjectId == 0 {
		base += float64(c.Stats.BasePDef.Underwear)
	}

	for _, v := range c.BonusStats {
		if v.Type == "physical_defense" {
			base += v.Val
		}
	}
	base *= float64(c.Level+89) / 100

	return int32(base)
}

func (c *Character) GetInventoryLimit() int16 {
	if c.Race == race.DWARF {
		return 100
	}
	return 80
}

func (c *Character) setWorldRegion(newRegion interfaces.WorldRegioner) {
	var oldAreas []interfaces.WorldRegioner

	currReg := c.GetCurrentRegion().(*WorldRegion)
	if currReg != nil {
		c.CurrentRegion.DeleteVisibleChar(c)
		oldAreas = currReg.GetNeighbors()
	}

	var newAreas []interfaces.WorldRegioner
	if newRegion != nil {
		newRegion.AddVisibleChar(c)
		newAreas = newRegion.GetNeighbors()
	}

	// кому отправить charInfo
	deleteObjectPkgTo := make([]int32, 0, 64)
	for _, region := range oldAreas {
		if !Contains(newAreas, region) {

			for _, v := range region.GetCharsInRegion() {
				if v.GetObjectId() == c.GetObjectId() {
					continue
				}
				deleteObjectPkgTo = append(deleteObjectPkgTo, v.GetObjectId())
			}
		}
	}
	if len(deleteObjectPkgTo) > 0 {
		c.DeleteObjectTo <- deleteObjectPkgTo
	}

	// кому отправить charInfo
	charInfoPkgTo := make([]int32, 0, 64)
	npcPkgTo := make([]interfaces.Npcer, 0, 64)
	for _, region := range newAreas {
		if !Contains(oldAreas, region) {
			for _, v := range region.GetCharsInRegion() {
				if v.GetObjectId() == c.GetObjectId() {
					continue
				}
				charInfoPkgTo = append(charInfoPkgTo, v.GetObjectId())
			}

			npcPkgTo = append(npcPkgTo, region.GetNpcInRegion()...)

		}
	}
	if len(charInfoPkgTo) > 0 {
		c.CharInfoTo <- charInfoPkgTo
	}
	c.CurrentRegion = newRegion.(*WorldRegion)

	if len(npcPkgTo) > 0 {
		c.NpcInfo <- npcPkgTo
	}

}

func (c *Character) checkRegion() {
	for {
		time.Sleep(time.Second)
		if c.CurrentRegion != nil {
			curReg := c.CurrentRegion
			x, y, z := c.GetXYZ()
			ncurReg := GetRegion(x, y, z)
			if curReg != ncurReg {
				c.setWorldRegion(ncurReg)
			}
		}
	}

}

// SaveFirstInGamePlayer Сохранение отметки что юзер зашел в игру впервый раз с момента создания игрока
func (c *Character) SaveFirstInGamePlayer() {
	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	sql := `UPDATE "characters" SET "first_enter_game" = false WHERE "object_id" = $1`
	_, err = dbConn.Exec(context.Background(), sql, c.ObjectId)
	if err != nil {
		logger.Error.Panicln(err)
	}
	c.FirstEnterGame = false
}

// ExistItemInInventory Возвращает ссылку на Item если он есть в инвентаре
func (c *Character) ExistItemInInventory(objectItemId int32) *MyItem {
	for i := range c.Inventory.Items {
		item := &c.Inventory.Items[i]
		if item.ObjectId == objectItemId {
			return item
		}
	}
	return nil
}

func (c *Character) GetObjectId() int32 {
	return c.ObjectId
}
func (c *Character) GetAccountLogin() string {
	return c.Login
}
func (c *Character) GetName() string {
	return c.CharName
}
func (c *Character) SetX(x int32) {
	c.Coordinates.X = x
}
func (c *Character) SetY(y int32) {
	c.Coordinates.Y = y
}
func (c *Character) SetZ(z int32) {
	c.Coordinates.Z = z
}
func (c *Character) SetXYZ(x, y, z int32) {
	c.Coordinates.X = x
	c.Coordinates.Y = y
	c.Coordinates.Z = z
}
func (c *Character) SetHeading(h int32) {
	c.Heading = h
}
func (c *Character) SetInstanceId(i int32) {
	_ = i
	//TODO release
}
func (c *Character) GetXYZ() (x, y, z int32) {
	return c.Coordinates.X, c.Coordinates.Y, c.Coordinates.Z
}
func (c *Character) GetX() int32 {
	return c.Coordinates.X
}
func (c *Character) GetY() int32 {
	return c.Coordinates.Y
}
func (c *Character) GetZ() int32 {
	return c.Coordinates.Z
}

func (c *Character) GetCurrentRegion() interfaces.WorldRegioner {
	return c.CurrentRegion
}

func (c *Character) CloseChannels() {
	c.ChannelUpdateShadowItem = nil
	c.NpcInfo = nil
	c.CharInfoTo = nil
	c.DeleteObjectTo = nil
}
func (c *Character) StartTransactionRequest() {
	c.RequestExpireTime = time.Now().Add(RequestTimeout).Unix()
}
func (c *Character) IsProcessingRequest() bool {
	return c.RequestExpireTime > time.Now().Unix()
}
func (c *Character) IsProcessingTransaction() bool {
	return c.ActiveTradeList != nil || c.RequestExpireTime > time.Now().Unix()
}

func (c *Character) GetClassId() int32 {
	return c.ClassId
}
func (c *Character) CalculateDistanceTo(ox, oy, oz int32, includeZAxis, squared bool) float64 {
	return CalculateDistance(c.GetX(), c.GetY(), c.GetZ(), ox, oy, oz, includeZAxis, squared)
}
func (c *Character) GetTradeRefusal() bool {
	return c.TradeRefusal
}
func (c *Character) OnTransactionRequest(p interfaces.CharacterI) {
	c.StartTransactionRequest()
	p.SetActiveRequester(c)
}

func (c *Character) SetActiveRequester(partner interfaces.CharacterI) {
	c.ActiveRequester = partner
}

func (c *Character) GetActiveRequester() interfaces.CharacterI {
	return c.ActiveRequester
}

func (c *Character) OnTransactionResponse() {
	c.RequestExpireTime = time.Now().Unix()
}

func (c *Character) StartTrade(partner interfaces.CharacterI) {
	c.OnTradeStart(partner)
	partner.OnTradeStart(c)
}

func (c *Character) OnTradeStart(partner interfaces.CharacterI) {
	c.ActiveTradeList = NewTradeList(c)
	c.ActiveTradeList.SetPartner(partner)

}

func (c *Character) IsRequestExpired() bool {
	return c.RequestExpireTime < time.Now().Unix()
}

func (c *Character) GetActiveTradeList() interfaces.TradeListInterface {
	return c.ActiveTradeList
}

func (c *Character) CloseConnection() {
	c.Conn.CloseConnection()
}

// CancelActiveTrade
// возвращает bool,bool.
// Надо ли отправлять tradeDone(0) и sysMsg для себя(первый параметр) и партнёра(второй параметр)
func (c *Character) CancelActiveTrade() (bool, bool) {
	if c.ActiveTradeList == nil {
		return false, false
	}
	needPartnerSendPacket := false
	partner := c.ActiveTradeList.GetPartner()
	if partner != nil {
		needPartnerSendPacket = partner.OnTradeCancel()
	}

	return c.OnTradeCancel(), needPartnerSendPacket
}

func (c *Character) OnTradeCancel() bool {
	if c.ActiveTradeList == nil {
		return false
	}
	c.ActiveTradeList.Lock()
	c.ActiveTradeList = nil
	return true
}

func (c *Character) OnTradeFinish() {
	c.ActiveTradeList = nil
	//c.EncryptAndSend(serverpackets.TradeDone(1))
	//if successful {
	//	c.EncryptAndSend(sysmsg.SystemMessage(sysmsg.TradeSuccessful))
	//}
}

func (c *Character) ValidateItemManipulation(objectId int32) bool {
	item := c.Inventory.GetItemByObjectId(objectId)
	if item == nil {
		return false
	}
	if c.ActiveEnchantItemId == objectId {
		return false
	}
	//todo доделть проверка на хаус оружие
	// проверка на пета
	return true
}

func (c *Character) CheckItemManipulation(objectId int32, count int64) interfaces.MyItemInterface {
	// todo куча проверок
	item := c.Inventory.GetItemByObjectId(objectId)
	if item == nil {
		return nil
	}
	return item

}
func (c *Character) GetItemByObjectId(objectId int32) interfaces.MyItemInterface {
	return c.Inventory.GetItemByObjectId(objectId)
}

func (c *Character) GetInventory() interfaces.InventoryInterface {
	return &c.Inventory
}

func (c *Character) ValidateWeight(weight int32) bool {
	return c.Inventory.TotalWeight+weight <= c.GetMaxLoad()
}

func (c *Character) GetMaxLoad() int32 {
	//todo calcStat
	return 69000 * 3
}
func (c *Character) GetActiveEnchantItemId() int32 {
	return c.ActiveEnchantItemId
}

func (c *Character) SendSysMsg(num interface{}, options ...string) {
	smsg := num.(sysmsg.SysMsg)

	c.EncryptAndSend(sysmsg.SystemMessage(smsg))
}

// методы для реализации ClientInterface, не нужно их заполнять
func (c *Character) SetLogin(login string)                                     { panic("нельзя") }
func (c *Character) RemoveCurrentChar()                                        { panic("нельзя") }
func (c *Character) SetState(state clientStates.State)                         { panic("нельзя") }
func (c *Character) GetState() clientStates.State                              { panic("нельзя") }
func (c *Character) SetSessionKey(playOk1, playOk2, loginOk1, loginOk2 uint32) { panic("нельзя") }
func (c *Character) GetSessionKey() (playOk1, playOk2, loginOk1, loginOk2 uint32) {
	panic("нельзя")
}

///////////

// методлы для реализации ReciverAndSender
func (c *Character) Receive() (opcode byte, data []byte, err error) { panic("нельзя") }
func (c *Character) AddLengthAndSand(data []byte)                   { c.Conn.AddLengthAndSand(data) }
func (c *Character) Send(data []byte)                               { c.Conn.Send(data) }
func (c *Character) SendBuf(buffer *packets.Buffer) error           { return c.Conn.SendBuf(buffer) }
func (c *Character) EncryptAndSend(data []byte)                     { c.Conn.EncryptAndSend(data) }
func (c *Character) CryptAndReturnPackageReadyToShip(data []byte) []byte {
	return c.Conn.CryptAndReturnPackageReadyToShip(data)
}
func (c *Character) GetCurrentChar() interfaces.CharacterI { return c }

///////////

func (c *Character) DropItem(objectId int32, count int64) interfaces.MyItemInterface {
	//invitem := c.Inventory.GetItemByObjectId(objectId)
	item := c.Inventory.DropItem(objectId, count)

	if item == nil {
		return nil
	}

	return item
}
