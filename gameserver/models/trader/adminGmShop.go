package trader

import (
	"encoding/json"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/items"
	"os"
)

var ADMIN_SHOP_FILE = "./datapack/data/trader/gmshop.json"

var globalGmShopManager *GmShopManager

type GmShopItem struct {
	ItemID int                        `json:"item_id"`
	Item   interfaces.MyItemInterface `json:"item,omitempty"`
}

// Реализация интерфейса TraderGmShopItem
func (gsi *GmShopItem) GetItemID() int {
	return gsi.ItemID
}

func (gsi *GmShopItem) GetItem() interfaces.MyItemInterface {
	return gsi.Item
}

type GmShopData struct {
	ShopID int          `json:"shop"`
	Items  []GmShopItem `json:"items"`
}

func (gsd *GmShopData) GetShopID() int {
	return gsd.ShopID
}

func (gsd *GmShopData) GetItems() []interfaces.TraderGmShopItem {
	shopItems := make([]interfaces.TraderGmShopItem, len(gsd.Items))
	for i := range gsd.Items {
		shopItems[i] = &gsd.Items[i]
	}
	return shopItems
}

type GmShopList []GmShopData

type GmShopManager struct {
	shops       GmShopList
	shopByID    map[int]*GmShopData
	shopByNpcID map[int]*GmShopData
}

func (gm *GmShopManager) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var rawShops []RawGmShopData
	if err := json.Unmarshal(data, &rawShops); err != nil {
		logger.LogError(err.Error())
		return err
	}

	var shops GmShopList
	shopByID := make(map[int]*GmShopData)

	for _, rawShop := range rawShops {
		var shopItems []GmShopItem
		for _, itemID := range rawShop.Items {
			itemData, ok := items.GetItemInfo(itemID)
			if !ok {
				logger.LogError("Ошибка загрузки itemID: %d", itemID)
				continue
			}
			myItem := &models.MyItem{Item: itemData}

			shopItems = append(shopItems, GmShopItem{
				ItemID: itemID,
				Item:   myItem,
			})
		}

		shop := GmShopData{
			ShopID: rawShop.ShopID,
			Items:  shopItems,
		}

		shops = append(shops, shop)
		shopByID[shop.ShopID] = &shops[len(shops)-1]
		gm.shopByNpcID[rawShop.NpcID] = &shops[len(shops)-1]
	}

	gm.shops = shops
	gm.shopByID = shopByID

	return nil
}

func (gm *GmShopManager) GetShopByID(shopID int) (interfaces.TraderGmShopData, bool) {
	shop := gm.shopByID[shopID]
	if shop == nil {
		return nil, false
	}
	return shop, true
}

func (gm *GmShopManager) GetShopsCount() int {
	return len(gm.shops)
}

func (gm *GmShopManager) ClearShops() {
	gm.shops = make(GmShopList, 0)
	gm.shopByID = make(map[int]*GmShopData)
}

func (gm *GmShopManager) ReloadShops() {
	gm.ClearShops()
	if err := gm.LoadFromFile(ADMIN_SHOP_FILE); err != nil {
		logger.LogError("Ошибка при перезагрузке магазинов: %v", err)
	}
}

type RawGmShopData struct {
	NpcID  int   `json:"npc"`
	ShopID int   `json:"shop"`
	Items  []int `json:"items"`
}

func NewGmShopManager() *GmShopManager {
	return &GmShopManager{
		shops:       make(GmShopList, 0),
		shopByID:    make(map[int]*GmShopData),
		shopByNpcID: make(map[int]*GmShopData),
	}
}

func LoadShops() {
	manager := NewGmShopManager()
	err := manager.LoadFromFile(ADMIN_SHOP_FILE)
	if err != nil {
		logger.LogError("Ошибка загрузки GM магазинов: %v", err)
		return
	}
	globalGmShopManager = manager
	logger.LogInfo("Загружено %d GM магазинов", len(manager.shops))
}

func GetGmShopManager() *GmShopManager {
	return globalGmShopManager
}
