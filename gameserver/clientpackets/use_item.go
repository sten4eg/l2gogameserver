package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/models/race"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

const FormalWearId = 6408

func UseItem(data []byte, client *models.Client) {
	var packet = packets.NewReader(data)

	objId := packet.ReadInt32() // targetObjId
	ctrlPressed := packet.ReadInt32() != 0
	_ = ctrlPressed

	var selectedItem models.MyItem

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
			if client.CurrentChar.IsActiveWeapon() && models.GetActiveWeapon(client.CurrentChar.Inventory, client.CurrentChar.Paperdoll).Item.Id == 9819 {
				serverpackets.SystemMessage(sysmsg.CannotEquipItemDueToBadCondition, client)
				return
			}
			//todo тут 2 проврки на  isMounted  и isDisarmed

			// нельзя менять оружие/щит если в руках проклятое оружие
			if client.CurrentChar.IsCursedWeaponEquipped() {
				return
			}

			//  запрет носить НЕ камаелям эксклюзивное оружие  камаелей
			if selectedItem.IsEquipped() == 0 && selectedItem.IsWeapon() { // todo еще проверка && !activeChar.canOverrideCond(ITEM_CONDITIONS))

				switch client.CurrentChar.Race {
				case race.KAMAEL:
					if selectedItem.IsWeaponTypeNone() {
						serverpackets.SystemMessage(sysmsg.CannotEquipItemDueToBadCondition, client)
						return
					}
				case race.HUMAN, race.DWARF, race.ELF, race.DARK_ELF, race.ORC:
					if selectedItem.IsOnlyKamaelWeapon() {
						serverpackets.SystemMessage(sysmsg.CannotEquipItemDueToBadCondition, client)
						return
					}
				}
			}
		// камаель не может носить тяжелую или маг броню
		// они могут носить только лайт, может проверять на !LIGHT ?
		case items.SlotChest, items.SlotBack, items.SlotGloves, items.SlotFeet, items.SlotHead, items.SlotFullArmor, items.SlotLegs:
			if client.CurrentChar.Race == race.KAMAEL && (selectedItem.IsHeavyArmor() || selectedItem.IsMagicArmor()) {
				serverpackets.SystemMessage(sysmsg.CannotEquipItemDueToBadCondition, client)
				return
			}
		case items.SlotDeco:
			//todo проверка !item.isEquipped() && (activeChar.getInventory().getTalismanSlots() == 0

		}

	}

	models.UseEquippableItem(selectedItem, client.CurrentChar)

	models.SaveInventoryInDB(client.CurrentChar.Inventory)

	serverpackets.InventoryUpdate(client, models.UpdateTypeModify)

	// После каждого use_item будет запрос в бд на восстановление paperdoll,
	// надо бы это сделать в UseEquippableItem
	client.CurrentChar.Paperdoll = models.RestoreVisibleInventory(client.CurrentChar.CharId)

	serverpackets.UserInfo(client)

	client.SentToSend()
}
