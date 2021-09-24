package models

import (
	"context"
	"l2gogameserver/db"
)

func (c *Client) SaveUser () {
	c.saveLocation()
}

func (c *Client) saveLocation(){
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}

	defer dbConn.Release()
	sql := `UPDATE "characters" SET "x" = $1, "y" = $2, "z" = $3 WHERE "id" = $4`
	x, y, z := c.CurrentChar.GetXYZ()
	_ , err = dbConn.Exec(context.Background(), sql, x, y, z, c.CurrentChar.CharId)
	if err != nil {
		panic(err)
	}
}