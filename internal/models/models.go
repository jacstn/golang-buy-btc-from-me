package models

import (
	"database/sql"
)

type Url struct {
	Id        uint16
	Name      string
	CreatedAt string
}

func SaveModel(db *sql.DB, url Url) int64 {
	return 0
}
