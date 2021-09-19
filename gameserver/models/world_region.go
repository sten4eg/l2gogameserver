package models

import (
	"sync"
)

type WorldRegion struct {
	TileX         int32
	TileY         int32
	Sur           []*WorldRegion
	CharsInRegion sync.Map //TODO переделать на мапу с RW мьютексом ци шо ци каво
}

func NewWorldRegion(x, y int32) WorldRegion {
	var newRegion WorldRegion
	newRegion.TileX = x
	newRegion.TileY = y
	return newRegion
}

func (w *WorldRegion) AddSurrounding(s *WorldRegion) {
	w.Sur = append(w.Sur, s)
}

func (w *WorldRegion) AddVisibleObject(character *Character) {
	w.CharsInRegion.Store(character.CharId, character)
}

func (w *WorldRegion) DeleteVisibleObject(character *Character) {
	w.CharsInRegion.Delete(character)
}

func GetAroundPlayersInRadius(me *Character, radius int32) []int32 {
	sqradius := int64(radius * radius)
	x, y, _ := me.GetXYZ()
	reg := GetRegion(x, y)
	var charIds []int32
	for _, region := range reg.Sur {
		region.CharsInRegion.Range(func(key, value interface{}) bool {
			val := value.(*Character)
			if val.CharId != me.CharId {
				dx := int64(val.Coordinates.X - x)
				dx *= dx
				if dx > sqradius {
					return true
				}
				dy := int64(val.Coordinates.Y - y)
				dy *= dy
				if dx+dy < sqradius {
					charIds = append(charIds, val.CharId)
				}
			}
			return true
		})
	}
	return charIds
}

func GetAroundCharacterInRadius(me *Character, radius int32) []*Character {
	sqradius := int64(radius * radius)
	x, y, _ := me.GetXYZ()
	reg := GetRegion(x, y)
	var charIds []*Character
	for _, region := range reg.Sur {
		region.CharsInRegion.Range(func(key, value interface{}) bool {
			val := value.(*Character)
			if val.CharId != me.CharId {
				dx := int64(val.Coordinates.X - x)
				dx *= dx
				if dx > sqradius {
					return true
				}
				dy := int64(val.Coordinates.Y - y)
				dy *= dy
				if dx+dy < sqradius {
					charIds = append(charIds, val)
				}
			}
			return true
		})
	}
	return charIds
}

func GetAroundPlayers(me *Character) []*Character {
	x, y, _ := me.GetXYZ()
	reg := GetRegion(x, y)
	var charIds []*Character
	for _, region := range reg.Sur {
		region.CharsInRegion.Range(func(key, value interface{}) bool {
			val := value.(*Character)
			if val.CharId != me.CharId {
				charIds = append(charIds, val)
			}
			return true
		})
	}
	return charIds
}
