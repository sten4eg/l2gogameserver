package clientpackets

import (
	"database/sql"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/models/race"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

const formalWearId = 6408
const fortFlagId = 9819

func UseItem(clientI interfaces.CharacterI, data []byte, db *sql.DB) {
	client, ok := clientI.(*models.Character)
	if !ok {
		return
	}
	var packet = packets.NewReader(data)

	objId := packet.ReadInt32() // targetObjId
	ctrlPressed := packet.ReadInt32() != 0
	_ = ctrlPressed

	var selectedItem *models.MyItem

	find := false
	for i := range client.Inventory.Items {
		item := &client.Inventory.Items[i]
		if item.ObjectId == objId {
			selectedItem = item
			find = true
			break
		}
	}

	// если предмет не найден в инвентаре, то выходим
	if !find {
		return
	}

	buffer := packets.Get()

	if selectedItem.IsEquipable() {
		// нельзя надевать Formal Wear с проклятым оружием
		if client.IsCursedWeaponEquipped() && objId == formalWearId {
			return
		}

		// todo тут еще 2 проверки

		switch selectedItem.SlotBitType {
		case items.SlotLrHand, items.SlotLHand, items.SlotRHand:

			// если в руке Combat flag
			if client.IsActiveWeapon() && models.GetActiveWeapon(client.Inventory.Items, client.Paperdoll).Item.Id == fortFlagId {
				pkg := sysmsg.SystemMessage(sysmsg.CannotEquipItemDueToBadCondition)
				buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
				client.Send(buffer.Bytes())
				return
			}
			//todo тут 2 проврки на  isMounted  и isDisarmed

			// нельзя менять оружие/щит если в руках проклятое оружие
			if client.IsCursedWeaponEquipped() {
				return
			}

			//  запрет носить НЕ камаелям эксклюзивное оружие  камаелей
			if selectedItem.IsEquipped() == 0 && selectedItem.IsWeapon() { // todo еще проверка && !activeChar.canOverrideCond(ITEM_CONDITIONS))

				switch client.Race {
				case race.KAMAEL:
					if selectedItem.IsWeaponTypeNone() {
						pkg := sysmsg.SystemMessage(sysmsg.CannotEquipItemDueToBadCondition)
						buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
						client.Send(buffer.Bytes())
						return
					}
				case race.HUMAN, race.DWARF, race.ELF, race.DARK_ELF, race.ORC:
					if selectedItem.IsOnlyKamaelWeapon() {
						pkg := sysmsg.SystemMessage(sysmsg.CannotEquipItemDueToBadCondition)
						buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
						client.Send(buffer.Bytes())
						return
					}
				}
			}
		// камаель не может носить тяжелую или маг броню
		// они могут носить только лайт, может проверять на !LIGHT ?
		case items.SlotChest, items.SlotBack, items.SlotGloves, items.SlotFeet, items.SlotHead, items.SlotFullArmor, items.SlotLegs:
			if client.Race == race.KAMAEL && (selectedItem.IsHeavyArmor() || selectedItem.IsMagicArmor()) {
				pkg := sysmsg.SystemMessage(sysmsg.CannotEquipItemDueToBadCondition)
				buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
				client.Send(buffer.Bytes())
				return
			}
		case items.SlotDeco:
			//todo проверка !item.isEquipped() && (activeChar.getInventory().getTalismanSlots() == 0

		}

		models.UseEquippableItem(selectedItem, client)
	}

	models.SaveInventoryInDB(client.Inventory.Items, db)

	pkg := serverpackets.InventoryUpdate(clientI.GetCurrentChar().GetInventory().GetItemsWithUpdatedType())
	clientI.GetCurrentChar().GetInventory().SetAllItemsUpdatedTypeNone()
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	// После каждого use_item будет запрос в бд на восстановление paperdoll,
	//todo надо бы это сделать в UseEquippableItem
	client.Paperdoll = models.RestoreVisibleInventory(client.ObjectId, db)

	pkg2 := serverpackets.UserInfo(client.GetCurrentChar())
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg2))

	client.Send(buffer.Bytes())
}
