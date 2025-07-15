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

func UseItem(clientI interfaces.CharacterI, data []byte, db *sql.DB, gs GsInterfNew) {
	character, ok := clientI.(*models.Character)
	if !ok {
		return
	}
	var packet = packets.NewReader(data)

	objId := packet.ReadInt32() // targetObjId
	ctrlPressed := packet.ReadInt32() != 0
	_ = ctrlPressed

	var selectedItem *models.PlayerItem

	find := false
	for i := range character.Inventory.Items {
		item := &character.Inventory.Items[i]
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

	if selectedItem.GetItemInfo().IsEquipable() {
		// нельзя надевать Formal Wear с проклятым оружием
		if character.IsCursedWeaponEquipped() && objId == formalWearId {
			return
		}

		// todo тут еще 2 проверки

		switch selectedItem.GetItemInfo().SlotBitType {
		case items.SlotLrHand, items.SlotLHand, items.SlotRHand:

			// если в руке Combat flag
			if character.IsActiveWeapon() && models.GetActiveWeapon(character.Inventory.Items, character.Paperdoll).GetItemInfo().GetId() == fortFlagId {
				pkg := sysmsg.SystemMessage(sysmsg.CannotEquipItemDueToBadCondition)
				buffer.WriteSlice(character.CryptAndReturnPackageReadyToShip(pkg))
				character.Send(buffer.Bytes())
				return
			}
			//todo тут 2 проврки на  isMounted  и isDisarmed

			// нельзя менять оружие/щит если в руках проклятое оружие
			if character.IsCursedWeaponEquipped() {
				return
			}

			//  запрет носить НЕ камаелям эксклюзивное оружие  камаелей
			if selectedItem.IsEquipped() == 0 && selectedItem.GetItemInfo().IsWeapon() { // todo еще проверка && !activeChar.canOverrideCond(ITEM_CONDITIONS))

				switch character.Race {
				case race.KAMAEL:
					if selectedItem.GetItemInfo().IsWeaponTypeNone() {
						pkg := sysmsg.SystemMessage(sysmsg.CannotEquipItemDueToBadCondition)
						buffer.WriteSlice(character.CryptAndReturnPackageReadyToShip(pkg))
						character.Send(buffer.Bytes())
						return
					}
				case race.HUMAN, race.DWARF, race.ELF, race.DARK_ELF, race.ORC:
					if selectedItem.GetItemInfo().IsOnlyKamaelWeapon() {
						pkg := sysmsg.SystemMessage(sysmsg.CannotEquipItemDueToBadCondition)
						buffer.WriteSlice(character.CryptAndReturnPackageReadyToShip(pkg))
						character.Send(buffer.Bytes())
						return
					}
				}
			}
		// камаель не может носить тяжелую или маг броню
		// они могут носить только лайт, может проверять на !LIGHT ?
		case items.SlotChest, items.SlotBack, items.SlotGloves, items.SlotFeet, items.SlotHead, items.SlotFullArmor, items.SlotLegs:
			if character.Race == race.KAMAEL && (selectedItem.GetItemInfo().IsHeavyArmor() || selectedItem.GetItemInfo().IsMagicArmor()) {
				pkg := sysmsg.SystemMessage(sysmsg.CannotEquipItemDueToBadCondition)
				buffer.WriteSlice(character.CryptAndReturnPackageReadyToShip(pkg))
				character.Send(buffer.Bytes())
				return
			}
		case items.SlotDeco:
			//todo проверка !item.isEquipped() && (activeChar.getInventory().getTalismanSlots() == 0

		}

		models.UseEquippableItem(selectedItem, character)
	}

	models.SaveInventoryInDB(character.Inventory.Items, db)

	pkg := serverpackets.InventoryUpdate(character.GetInventory().GetItemsWithUpdatedType())
	character.GetInventory().SetAllItemsUpdatedTypeNone()
	buffer.WriteSlice(character.CryptAndReturnPackageReadyToShip(pkg))

	// После каждого use_item будет запрос в бд на восстановление paperdoll,
	//todo надо бы это сделать в UseEquippableItem
	character.Paperdoll = models.RestoreVisibleInventoryWithCharacter(character, db)

	pkg2 := serverpackets.UserInfo(character.GetCurrentChar())
	buffer.WriteSlice(character.CryptAndReturnPackageReadyToShip(pkg2))

	client := gs.GetClientByLogin(character.GetAccountLogin())
	client.Send(buffer.Bytes())
}
