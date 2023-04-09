package db

import (
	"ariga.io/entcache"
	"code-connect/pkg/log"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	_ "github.com/mattn/go-sqlite3"
)

var l = log.NewZapSugaredLogger().With("pkg", "db")

func NewCachedEntDriver() *entcache.Driver {
	db, err := sql.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		l.Fatalw("failed opening connection to sqlite", "err", err)
	}
	// Decorates the sql.Driver with entcache.Driver.
	drv := entcache.NewDriver(db)

	return drv
}
