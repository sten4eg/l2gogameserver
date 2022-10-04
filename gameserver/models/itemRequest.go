package models

type ItemRequest struct {
	objectId int32
	itemId   int32
	count    int64
	price    int64
}

func NewItemRequest(objectId, itemId int32, count, price int64) *ItemRequest {
	var item ItemRequest
	item.objectId = objectId
	item.itemId = itemId
	item.count = count
	item.price = price
	return &item
}

func NewItemRequestWithoutItemId(objectId int32, count, price int64) *ItemRequest {
	var item ItemRequest
	item.objectId = objectId
	item.count = count
	item.price = price
	return &item
}

func (i *ItemRequest) GetObjectId() int32 {
	return i.objectId
}

func (i *ItemRequest) GetId() int32 {
	return i.itemId
}

func (i *ItemRequest) SetCount(count int64) {
	i.count = count
}

func (i *ItemRequest) GetCount() int64 {
	return i.count
}

func (i *ItemRequest) GetPrice() int64 {
	return i.price
}
