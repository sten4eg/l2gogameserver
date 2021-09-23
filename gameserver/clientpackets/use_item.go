package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/models/race"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

const formalWearId = 6408
const fortFlagId = 9819

func UseItem(client *models.Client, data []byte) []byte {
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
		return []byte{}
	}

	buffer := packets.Get()
	defer packets.Put(buffer)

	if selectedItem.IsEquipable() {
		// нельзя надевать Formal Wear с проклятым оружием
		if client.CurrentChar.IsCursedWeaponEquipped() && objId == formalWearId {
			return []byte{}
		}

		// todo тут еще 2 проверки

		switch selectedItem.SlotBitType {
		case items.SlotLrHand, items.SlotLHand, items.SlotRHand:

			// если в руке Combat flag
			if client.CurrentChar.IsActiveWeapon() && models.GetActiveWeapon(client.CurrentChar.Inventory, client.CurrentChar.Paperdoll).Item.Id == fortFlagId {
				pkg := serverpackets.SystemMessage(sysmsg.CannotEquipItemDueToBadCondition, client)
				buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
				return buffer.Bytes()
			}
			//todo тут 2 проврки на  isMounted  и isDisarmed

			// нельзя менять оружие/щит если в руках проклятое оружие
			if client.CurrentChar.IsCursedWeaponEquipped() {
				return []byte{}
			}

			//  запрет носить НЕ камаелям эксклюзивное оружие  камаелей
			if selectedItem.IsEquipped() == 0 && selectedItem.IsWeapon() { // todo еще проверка && !activeChar.canOverrideCond(ITEM_CONDITIONS))

				switch client.CurrentChar.Race {
				case race.KAMAEL:
					if selectedItem.IsWeaponTypeNone() {
						pkg := serverpackets.SystemMessage(sysmsg.CannotEquipItemDueToBadCondition, client)
						buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
						return buffer.Bytes()
					}
				case race.HUMAN, race.DWARF, race.ELF, race.DARK_ELF, race.ORC:
					if selectedItem.IsOnlyKamaelWeapon() {
						pkg := serverpackets.SystemMessage(sysmsg.CannotEquipItemDueToBadCondition, client)
						buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
						return buffer.Bytes()
					}
				}
			}
		// камаель не может носить тяжелую или маг броню
		// они могут носить только лайт, может проверять на !LIGHT ?
		case items.SlotChest, items.SlotBack, items.SlotGloves, items.SlotFeet, items.SlotHead, items.SlotFullArmor, items.SlotLegs:
			if client.CurrentChar.Race == race.KAMAEL && (selectedItem.IsHeavyArmor() || selectedItem.IsMagicArmor()) {
				pkg := serverpackets.SystemMessage(sysmsg.CannotEquipItemDueToBadCondition, client)
				buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
				return buffer.Bytes()
			}
		case items.SlotDeco:
			//todo проверка !item.isEquipped() && (activeChar.getInventory().getTalismanSlots() == 0

		}

		models.UseEquippableItem(selectedItem, client.CurrentChar)

	}

	models.SaveInventoryInDB(client.CurrentChar.Inventory)

	pkg := serverpackets.InventoryUpdate(client, selectedItem.ObjId, models.UpdateTypeModify)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	// После каждого use_item будет запрос в бд на восстановление paperdoll,
	//todo надо бы это сделать в UseEquippableItem
	client.CurrentChar.Paperdoll = models.RestoreVisibleInventory(client.CurrentChar.ObjectId)

	pkg2 := serverpackets.UserInfo(client)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg2))

	return buffer.Bytes()
}
