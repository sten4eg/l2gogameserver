package interfaces

type TraderGmShopItem interface {
	GetItemID() int
	GetItem() MyItemInterface
}

type TraderGmShopData interface {
	GetShopID() int
	GetItems() []TraderGmShopItem
}

type TraderGmShopManager interface {
	LoadFromFile(filename string) error
	GetShopByID(shopID int) TraderGmShopData
	GetShopsCount() int
	ClearShops()
	ReloadShops()
}
