package interfaces

import (
	"database/sql"

	"l2gogameserver/gameserver/models/items/attribute"
	"l2gogameserver/packets"
	"sync"
)

type MyItemInterface interface {
	sync.Locker
	BaseItemInterface
	UniquerId
	TradableItemInterface
	IsEquipped() int16
	GetAttackElementType() attribute.Attribute
	GetCount() int64
	GetEnchant() int16
	GetLocation() string
	GetAttackElementPower() int16
	GetElementDefAttr() [6]int16
	GetEnchantedOption() [3]int32
	GetUpdateType() int16
	GetLocData() int32
	SetLocData(int32)
	GetMana() int32
	ChangeCount(int)
	SetUpdateType(int16)
	SetCount(int64)
	UpdateDB(*sql.DB)
	GetOwnerId() int32
	SetOwnerId(ownerId int32)
	SetCoordinate(x, y, z int32)
	GetCoordinate() (x, y, z int32)
	GetDefaultPrice() int
	GetTime() int
	SetCharacter(CharacterI)
	GetCharacter() CharacterI
	UseEquippableItem()
}

type ItemRequestInterface interface {
	UniquerId
	Identifier
	SetCount(int64)
	GetCount() int64
	GetPrice() int64
}

type BaseItemInterface interface {
	Identifier
	IsEquipable() bool
	IsHeavyArmor() bool
	IsMagicArmor() bool
	IsArmor() bool
	IsOnlyKamaelWeapon() bool
	IsWeapon() bool
	IsWeaponTypeNone() bool
	IsStackable() bool
	GetBaseItem() BaseItemInterface
	GetItemType1() int
	GetItemType2() int16
	GetWeight() int
	GetSlotBitType() int32
}

type TradableItemInterface interface {
	UniquerId
	Identifier
	BaseItemInterface
	GetBodyPart() int32
	GetEnchant() int16
	GetAttackElementType() attribute.Attribute
	GetAttackElementPower() int16
	GetElementDefAttr() [6]int16
	GetEnchantedOption() [3]int32
	GetCount() int64
	SetCount(count int64)
	GetLocData() int32
	IsEquipped() int16
	GetDefaultPrice() int
	GetPrice() int64
	WriteItem(buffer *packets.Buffer)
	SetStoreCount(int64)
	GetStoreCount() int64
	SetPrice(int64)
	SetObjectId(int32)
	SetEnchant(int16)
	GetMana() int32
	GetTime() int
}
