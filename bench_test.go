package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/alphadose/haxmap"
	"github.com/puzpuzpuz/xsync/v4"
)

const (
	numKeys      = 100000 // Количество элементов в мапе
	updatePeriod = time.Millisecond * 10
)

func BenchmarkHaxMap(b *testing.B) {
	m := haxmap.New[int32, int64]()

	// Инициализация мапы
	for i := 0; i < numKeys; i++ {
		m.Set(int32(i), time.Now().Unix())
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := int32(r.Intn(numKeys))
			m.Set(key, time.Now().Unix()+int64(r.Intn(60)))
			time.Sleep(updatePeriod)
		}

	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now().Unix()
		m.ForEach(func(key int32, value int64) bool {
			// Здесь происходит сравнение
			if value < now {
				m.Del(key)
			}
			return true
		})
	}
}

func BenchmarkXsyncMap(b *testing.B) {
	m := xsync.NewMap[int32, int64]()

	// Инициализация мапы
	for i := 0; i < numKeys; i++ {
		m.Store(int32(i), time.Now().Unix())
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	b.RunParallel(func(pb *testing.PB) {
		key := int32(r.Intn(numKeys))
		m.Store(key, time.Now().Unix()+int64(r.Intn(60)))
		time.Sleep(updatePeriod)
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.Now().Unix()
		m.Range(func(key int32, value int64) bool {
			if value < now {
				m.Delete(key)
			}
			return true
		})

	}
}
