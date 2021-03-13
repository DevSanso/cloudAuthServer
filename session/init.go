package session

import (
	"github.com/tidwall/buntdb"
)

var _db *buntdb.DB

func Init() error {
	if _db != nil {panic("already init session db")}
	db,err := buntdb.Open(":memory:")
	if err != nil {return err}
	_db = db

	return err
}

func Close() error {
	if _db == nil {panic("nil pointer db")}
	var temp = _db
	_db = nil
	return temp.Close()
}