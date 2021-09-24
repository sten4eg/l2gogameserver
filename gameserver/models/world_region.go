package models

import (
	"math"
	"sync"
)

type WorldRegion struct {
	TileX         int32
	TileY         int32
	TileZ         int32
	CharsInRegion sync.Map //TODO переделать на мапу с RW мьютексом ци шо ци каво
	NpcInRegion   sync.Map //TODO переделать на мапу с RW мьютексом ци шо ци каво
}

func NewWorldRegion(x, y, z int32) *WorldRegion {
	var newRegion WorldRegion
	newRegion.TileX = x
	newRegion.TileY = y
	newRegion.TileZ = z
	return &newRegion
}

func (w *WorldRegion) AddVisibleChar(character *Character) {
	w.CharsInRegion.Store(character.CharId, character)
}
func (w *WorldRegion) DeleteVisibleChar(character *Character) {
	w.CharsInRegion.Delete(character)
}

func (w *WorldRegion) AddVisibleNpc(npc Npc) {
	w.NpcInRegion.Store(npc.ObjId, npc)
}

func (w *WorldRegion) getNeighbors() []*WorldRegion {
	return GetNeighbors(int(w.TileX), int(w.TileY), int(w.TileZ), 1, 1)
}

func Contains(regions []*WorldRegion, region *WorldRegion) bool {
	for _, v := range regions {
		if v == region {
			return true
		}
	}
	return false
}
func GetAroundPlayer(c *Character) []*Character {
	currentRegion := c.CurrentRegion
	if nil == currentRegion {
		return nil
	}
	result := make([]*Character, 0, 64)

	for _, v := range currentRegion.getNeighbors() {
		v.CharsInRegion.Range(func(key, value interface{}) bool {
			result = append(result, value.(*Character))
			return true
		})
	}
	return result
}

func GetAroundPlayersInRadius(c *Character, radius int32) []*Character {
	currentRegion := c.CurrentRegion
	if nil == currentRegion {
		return nil
	}
	result := make([]*Character, 0, 64)

	sqrad := radius * radius

	for _, v := range currentRegion.getNeighbors() {
		v.CharsInRegion.Range(func(key, value interface{}) bool {
			char := value.(*Character)
			if char.CharId == c.CharId {
				return true
			}

			dx := math.Abs(float64(char.Coordinates.X - c.Coordinates.X))
			if dx > float64(radius) {
				return true
			}

			dy := math.Abs(float64(char.Coordinates.X - c.Coordinates.X))
			if dy > float64(radius) {
				return true
			}

			if dx*dx+dy*dy > float64(sqrad) {
				return true
			}

			result = append(result, char)
			return true
		})
	}
	return result
}
