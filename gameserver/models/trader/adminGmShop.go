package trader

import (
	"encoding/json"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/models/items"
	"os"
)

var ADMIN_SHOP_FILE = "./datapack/data/trader/gmshop.json"

var globalGmShopManager *GmShopManager

type GmShopItem struct {
	ItemID int        `json:"item_id"`
	Item   items.Item `json:"item,omitempty"`
}

type GmShopData struct {
	ShopID int          `json:"shop"`
	Items  []GmShopItem `json:"items"`
}

type GmShopList []GmShopData

type GmShopManager struct {
	shops       GmShopList
	shopByID    map[int]*GmShopData
	shopByNpcID map[int]*GmShopData
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
			shopItems = append(shopItems, GmShopItem{
				ItemID: itemID,
				Item:   *itemData,
			})
		}

		shop := GmShopData{
			ShopID: rawShop.ShopID,
			Items:  shopItems,
		}

		shops = append(shops, shop)
		shopByID[shop.ShopID] = &shops[len(shops)-1]
	}

	gm.shops = shops
	gm.shopByID = shopByID

	return nil
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

func (gm *GmShopManager) GetShopByID(shopID int) *GmShopData {
	return gm.shopByID[shopID]
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
	_ = gm.LoadFromFile(ADMIN_SHOP_FILE)
}

// Обёртки

func GetShopByID(shopID int) (*GmShopData, bool) {
	if globalGmShopManager == nil {
		return nil, false
	}
	shop := globalGmShopManager.GetShopByID(shopID)
	return shop, shop != nil
}

func GetShopsCount() int {
	if globalGmShopManager == nil {
		return 0
	}
	return globalGmShopManager.GetShopsCount()
}
