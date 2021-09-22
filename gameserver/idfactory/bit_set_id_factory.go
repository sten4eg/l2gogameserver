package idfactory

import (
	"context"
	"github.com/bits-and-blooms/bitset"
	"l2gogameserver/db"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
)

var IdExtracts = [][]string{
	{"characters", "id"},
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
	FreeIds = *bitset.New(100000)
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
}

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
			//increaseBitSetCapacity()
		} else {
			panic("Закончились objectId")
		}
	}
	atomic.StoreUint64(&NextFreeId, uint64(nextFree))

	mu.Unlock()
	return int32(newID + FirstOid)
}

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
