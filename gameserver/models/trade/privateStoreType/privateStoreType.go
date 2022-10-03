package privateStoreType

type PrivateStoreType byte

const (
	NONE         PrivateStoreType = 0
	SELL         PrivateStoreType = 1
	SELL_MANAGE  PrivateStoreType = 2
	BUY          PrivateStoreType = 3
	BUY_MANAGE   PrivateStoreType = 4
	MANUFACTURE  PrivateStoreType = 5
	PACKAGE_SELL PrivateStoreType = 8
)
