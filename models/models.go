package models

import "database/sql"

type StocksDB struct {
	*sql.DB
}
