package session

import (
	"github.com/tidwall/buntdb"
	"time"
)
func Set(key,value string) error {
	tx,err := _db.Begin(true)
	if err != nil {return err}
	
	_,_,err=tx.Set(key,value,nil)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func SetTimeOut(key,value string,duration time.Duration) error {
	tx,err := _db.Begin(true)
	if err != nil {return err}
	
	_,_,err=tx.Set(key,value,&buntdb.SetOptions{true,duration})
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}