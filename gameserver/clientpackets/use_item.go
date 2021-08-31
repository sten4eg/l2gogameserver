package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/models/items/armorType"
	"l2gogameserver/gameserver/models/items/weaponType"
	"l2gogameserver/gameserver/models/race"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

const FormalWearId = 6408

func NewUseItem(data []byte, client *models.Client) {
	var packet = packets.NewReader(data)

	objId := packet.ReadInt32() // targetObjId
	ctrlPressed := packet.ReadInt32() != 0
	_ = ctrlPressed

	var selectedItem items.MyItem

	find := false
	for _, v := range client.CurrentChar.Inventory {
		if v.ObjId == objId {
			selectedItem = v
			find = true
			break
		}
	}

	// если предмет не найден в инвентаре, то выходим
	if !find {
		return
	}

	if selectedItem.IsEquipable() {
		// нельзя надевать Formal Wear с проклятым оружием
		if client.CurrentChar.IsCursedWeaponEquipped() && objId == FormalWearId {
			return
		}

		// todo тут еще 2 проверки

		switch selectedItem.SlotBitType {
		case items.SlotLrHand, items.SlotLHand, items.SlotRHand:

			// если в руке Combat flag
			if client.CurrentChar.IsActiveWeapon() && items.GetActiveWeapon(client.CurrentChar.Inventory, client.CurrentChar.Paperdoll).Item.Id == 9819 {
				serverpackets.NewSystemMessage(sysmsg.CannotEquipItemDueToBadCondition, client)
				return
			}
			//todo тут 2 проврки на  isMounted  и isDisarmed

			// нельзя менять оружие/щит если в руках проклятое оружие
			if client.CurrentChar.IsCursedWeaponEquipped() {
				return
			}

			//  запрет носить не камаелям эксклюзивное оружие  камаелей
			if selectedItem.IsEquipped() == 0 && selectedItem.ItemType == items.Weapon { // todo еще проверка && !activeChar.canOverrideCond(ITEM_CONDITIONS))

				switch client.CurrentChar.Race {
				case race.KAMAEL:
					if selectedItem.WeaponType == weaponType.NONE {
						serverpackets.NewSystemMessage(sysmsg.CannotEquipItemDueToBadCondition, client)
						return
					}
				case race.HUMAN, race.DWARF, race.ELF, race.DARK_ELF, race.ORC:
					if selectedItem.WeaponType == weaponType.RAPIER || selectedItem.WeaponType == weaponType.CROSSBOW || selectedItem.WeaponType == weaponType.ANCIENTSWORD {
						serverpackets.NewSystemMessage(sysmsg.CannotEquipItemDueToBadCondition, client)
						return
					}
				}

			}
		// камаель не может носить тяжелую или маг броню
		// они могут носить только лайт, может проверять на !LIGHT ?
		case items.SlotChest, items.SlotBack, items.SlotGloves, items.SlotFeet, items.SlotHead, items.SlotFullArmor, items.SlotLegs:
			if client.CurrentChar.Race == race.KAMAEL && (selectedItem.ArmorType == armorType.HEAVY || selectedItem.ArmorType == armorType.MAGIC) {
				serverpackets.NewSystemMessage(sysmsg.CannotEquipItemDueToBadCondition, client)
				return
			}
		case items.SlotDeco:
			//todo проверка !item.isEquipped() && (activeChar.getInventory().getTalismanSlots() == 0

		}

	}

	if selectedItem.IsEquipped() == 1 {
		unEquipAndRecord(selectedItem, client.CurrentChar.Inventory)
	} else {
		equipItemAndRecord(selectedItem, client.CurrentChar.Inventory)
	}

	items.SaveInventoryInDB(client.CurrentChar.Inventory)

	serverpackets.NewInventoryUpdate(client, items.UpdateTypeModify)

	client.CurrentChar.Paperdoll = items.RestoreVisibleInventory(client.CurrentChar.CharId)

	serverpackets.NewUserInfo(client)

	client.SentToSend()
}

func unEquipAndRecord(item items.MyItem, myItems []items.MyItem) {
	switch item.SlotBitType {
	case items.SlotRHand:
		setPaperdollItem(items.PAPERDOLL_RHAND, nil, myItems)
	case items.SlotLegs:
		setPaperdollItem(items.PAPERDOLL_LEGS, nil, myItems)
	}
}

// equipItemAndRecord
func equipItemAndRecord(item items.MyItem, myItems []items.MyItem) {
	switch item.SlotBitType {
	case items.SlotRHand:
		setPaperdollItem(items.PAPERDOLL_RHAND, &item, myItems)
	case items.SlotLegs:
		setPaperdollItem(items.PAPERDOLL_LEGS, &item, myItems)
	}
}

func setPaperdollItem(slot uint8, item *items.MyItem, myItems []items.MyItem) {

	if item == nil {
		for i, v := range myItems {
			if v.LocData == int32(slot) {
				v.LocData = getFirstEmptySlot(myItems)
				v.Loc = items.Inventory
				myItems[i] = v
				break
			}
		}
		return
	}

	var old items.MyItem
	var k int
	var keyCurrentItem int
	for i, v := range myItems {
		if v.LocData == int32(slot) && v.Loc == items.Paperdoll { // todo if locdata or slot == 0
			k = i
			old = v
		}

		if v == *item {
			keyCurrentItem = i
		}

	}

	if old.Id != 0 {
		old.Loc = items.Inventory
		old.LocData = item.LocData
		myItems[k] = old
		item.LocData = int32(slot)
		item.Loc = items.Paperdoll
	} else {
		item.LocData = int32(slot)
		item.Loc = items.Paperdoll
	}

	myItems[keyCurrentItem] = *item
}

func getFirstEmptySlot(myItems []items.MyItem) int32 {
	var max int32
	for _, v := range myItems {
		if v.LocData > max {
			max = v.LocData
		}
	}

	var i int32
	for i = 0; i < max; i++ {
		flag := false
		for _, q := range myItems {
			if q.LocData == i && q.Loc != items.Paperdoll {
				flag = true
			}
		}

		if !flag {
			return i
		}
	}

	return max + 1
}
