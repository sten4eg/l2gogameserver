package models

import (
	"l2gogameserver/data/logger"
)

func (c *ClientCtx) SaveUser() {
	c.saveLocation()
}

func (c *ClientCtx) saveLocation() {
	sql := `UPDATE "characters" SET "x" = $1, "y" = $2, "z" = $3 WHERE "object_id" = $4`
	x, y, z := c.CurrentChar.GetXYZ()
	_, err := c.db.Exec(sql, x, y, z, c.CurrentChar.ObjectId)
	if err != nil {
		logger.Error.Panicln(err)
	}
}
