package models

import (
	"context"
	"l2gogameserver/data/logger"
	"l2gogameserver/db"
)

func (c *ClientCtx) SaveUser() {
	c.saveLocation()
}

func (c *ClientCtx) saveLocation() {
	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	sql := `UPDATE "characters" SET "x" = $1, "y" = $2, "z" = $3 WHERE "object_id" = $4`
	x, y, z := c.CurrentChar.GetXYZ()
	_, err = dbConn.Exec(context.Background(), sql, x, y, z, c.CurrentChar.ObjectId)
	if err != nil {
		logger.Error.Panicln(err)
	}
}
