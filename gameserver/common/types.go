package common

import (
	"sync"
)

// Coordinates структура для хранения координат
type Coordinates struct {
	mu sync.Mutex
	X  int32
	Y  int32
	Z  int32
}

// GetX возвращает X координату
func (c *Coordinates) GetX() int32 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.X
}

// GetY возвращает Y координату
func (c *Coordinates) GetY() int32 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Y
}

// GetZ возвращает Z координату
func (c *Coordinates) GetZ() int32 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Z
}

// SetX устанавливает X координату
func (c *Coordinates) SetX(x int32) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.X = x
}

// SetY устанавливает Y координату
func (c *Coordinates) SetY(y int32) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Y = y
}

// SetZ устанавливает Z координату
func (c *Coordinates) SetZ(z int32) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Z = z
}

// SetXYZ устанавливает все координаты
func (c *Coordinates) SetXYZ(x, y, z int32) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.X = x
	c.Y = y
	c.Z = z
}

// GetXYZ возвращает все координаты
func (c *Coordinates) GetXYZ() (x, y, z int32) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.X, c.Y, c.Z
}
