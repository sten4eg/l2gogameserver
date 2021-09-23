package idfactory

import (
	"context"
	"github.com/bits-and-blooms/bitset"
	"l2gogameserver/db"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var IdExtracts = [][]string{
	{"characters", "object_id"},
	{"items", "object_id"},
	//{"clan_data","clan_id"},//todo когда появятся таблиы - расскоментить
	//{"itemsonground","object_id"},
	//{"messages","messageId"},
}

var FreeIds bitset.BitSet
var FreeIdCount int32
var NextFreeId uint64

const FirstOid = 0x10000000
const LastOid = 0x7FFFFFFF
const FreeObjectIdSize = LastOid - FirstOid

var mu sync.Mutex

func Load() {
	primeInit()
	FreeIds = *bitset.New(uint(NextPrime(100000)))
	FreeIds.ClearAll()
	FreeIdCount = FreeObjectIdSize

	s, _ := FreeIds.NextClear(0)
	NextFreeId = uint64(s)

	for _, usedObjectId := range extractUsedObjectIDTable() {
		objectId := usedObjectId - FirstOid
		if objectId < 0 {
			panic("objectId меньше нуля")
		}

		FreeIds.Set(uint(objectId))
		atomic.AddInt32(&FreeIdCount, -1)
	}
	v, _ := FreeIds.NextClear(0)
	NextFreeId = uint64(v)

	go bitSetCapacityCheck()
}

// GetNext получение свободного идентификатора objectId
func GetNext() int32 {
	mu.Lock()
	newID := atomic.LoadUint64(&NextFreeId)
	FreeIds.Set(uint(newID))
	atomic.AddInt32(&FreeIdCount, -1)

	nextFree, _ := FreeIds.NextClear(uint(newID))
	if nextFree < 0 {
		nextFree, _ = FreeIds.NextClear(0)
	}

	if nextFree < 0 {
		if FreeIds.Len() < FreeObjectIdSize {
			increaseBitSetCapacity()
		} else {
			panic("Закончились objectId")
		}
	}
	atomic.StoreUint64(&NextFreeId, uint64(nextFree))

	mu.Unlock()
	return int32(newID + FirstOid)
}

// extractUsedObjectIDTable чтение из БД всех objectId
// и установка их как занятых
func extractUsedObjectIDTable() []int {
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err.Error())
	}
	defer dbConn.Release()

	sqlQuery := ""
	for _, column := range IdExtracts {
		sqlQuery += "SELECT " + column[1] + " FROM " + column[0] + " UNION "
	}
	sqlQuery = strings.TrimRight(sqlQuery, " UNION ")

	rows, err := dbConn.Query(context.Background(), sqlQuery)
	if err != nil {
		panic(err.Error())
	}
	var tmp []int
	for rows.Next() {
		var t int
		err = rows.Scan(&t)
		if err != nil {
			panic(err)
		}
		tmp = append(tmp, t)
	}

	sort.Ints(tmp)
	return tmp
}

// Release если objectId уже не используется его можно вернуть в пулл
// Чтобы в дальнейшем использовать снова
func Release(objectId int32) {
	mu.Lock()
	id := objectId - FirstOid
	if id > -1 {
		FreeIds.Clear(uint(id))
		atomic.AddInt32(&FreeIdCount, 1)
	} else {
		panic("Попытка release objectId")
	}
}

// usedIdCount количество использованных идентификаторов
func usedIdCount() int32 {
	return FreeIdCount - FirstOid
}

// increaseBitSetCapacity увеличение емкости BitSet
func increaseBitSetCapacity() {
	mu.Lock()
	newBitSet := bitset.New(uint(NextPrime(int(usedIdCount() * 11 / 10))))
	newBitSet.Union(&FreeIds)
	FreeIds = *newBitSet
	mu.Unlock()
}

// reachingBitSetCapacity достиг ли bitSet максимальной capacity
func reachingBitSetCapacity() bool {
	return uint(NextPrime(int(usedIdCount()*11/10))) > FreeIds.Len()
}

// bitSetCapacityCheck проверка каждые 30 секунд
// достиг ли bitSet максимальной capacity
// и увеличение его если необходимо
func bitSetCapacityCheck() {
	for {
		time.Sleep(time.Second * 30)
		if reachingBitSetCapacity() {
			increaseBitSetCapacity()
		}
	}
}
