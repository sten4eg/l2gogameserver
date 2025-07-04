package migrations

import "embed"

var (
	//go:embed *.sql
	FS embed.FS
)
