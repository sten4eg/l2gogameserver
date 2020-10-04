package models

import "sync"

type WorldRegion struct {
	TileX         int32
	TileY         int32
	Sur           []*WorldRegion
	CharsInRegion sync.Map
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
